package httpx

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type DoFunc func(*http.Client) (*http.Response, error)

type ClientConfig struct {
	Retry   int64
	Timeout time.Duration

	// Transport
	Proxy                 string
	DialerTimeout         time.Duration
	DialerKeepAlive       time.Duration
	ForceAttemptHTTP2     bool
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
}

type Client struct {
	Config ClientConfig
	*http.Client
}

type DumpData struct {
	Request               string
	DumpRequestError      string
	RequestBody           string
	DumpRequestBodyError  string
	Response              string
	DumpResponseError     string
	ResponseBody          string
	DumpResponseBodyError string
}

var DefaultClientConfig = ClientConfig{
	Retry:   1,
	Timeout: 60 * time.Second,

	// Transport
	Proxy:                 "",
	DialerTimeout:         30 * time.Second,
	DialerKeepAlive:       30 * time.Second,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          0,
	MaxIdleConnsPerHost:   2,
	MaxConnsPerHost:       0,
	TLSHandshakeTimeout:   10 * time.Second,
	ResponseHeaderTimeout: 10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	IdleConnTimeout:       180 * time.Second,
}

func Default() *Client {
	return NewClient(DefaultClientConfig)
}

func NewClient(config ClientConfig) *Client {
	c := &Client{
		Config: config,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: nil,
				DialContext: (&net.Dialer{
					Timeout:   config.DialerTimeout,
					KeepAlive: config.DialerKeepAlive,
				}).DialContext,
				ForceAttemptHTTP2:     config.ForceAttemptHTTP2,
				MaxIdleConns:          config.MaxIdleConns,
				MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
				MaxConnsPerHost:       config.MaxConnsPerHost,
				TLSHandshakeTimeout:   config.TLSHandshakeTimeout,
				ResponseHeaderTimeout: config.ResponseHeaderTimeout,
				ExpectContinueTimeout: config.ExpectContinueTimeout,
				IdleConnTimeout:       config.IdleConnTimeout,
			},
		},
	}

	if c.Config.Proxy != "" {
		proxyURL, err := url.Parse(c.Config.Proxy)
		if err != nil {
			log.Panicln(err)
		}

		c.Client.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	}

	c.Client.Timeout = c.Config.Timeout

	return c
}

func Request(doFunc DoFunc) ([]byte, error) {
	return Default().Request(doFunc)
}

func (c *Client) Request(doFunc DoFunc) (data []byte, err error) {
	var times int64
	for times < c.Config.Retry {
		times++
		data, err = c.do(doFunc)
		if err == nil {
			return
		}
	}

	return
}

func RequestWithDump(doFunc DoFunc) ([]byte, *DumpData, error) {
	return Default().RequestWithDump(doFunc)
}

func (c *Client) RequestWithDump(doFunc DoFunc) (data []byte, dumpData *DumpData, err error) {
	var times int64
	for times < c.Config.Retry {
		times++
		data, dumpData, err = c.doDump(doFunc)
		if err == nil {
			return
		}
	}

	return
}

func (c *Client) do(doFunc DoFunc) (responseBody []byte, err error) {
	response, doFuncErr := doFunc(c.Client)
	defer func() {
		if response != nil {
			response.Body.Close()
		}
	}()

	if response != nil {
		var readAllErr error
		responseBody, readAllErr = io.ReadAll(response.Body)
		if readAllErr != nil {
			err = readAllErr
			return nil, err
		}
	}

	if doFuncErr != nil {
		err = doFuncErr
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("status code of http response is %d instead of 200", response.StatusCode)
		return
	}

	return
}

func (c *Client) doDump(doFunc DoFunc) (responseBody []byte, dumpData *DumpData, err error) {
	response, doFuncErr := doFunc(c.Client)
	dumpData = c.dump(response)
	responseBody = []byte(dumpData.ResponseBody)

	if doFuncErr != nil {
		err = doFuncErr
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("status code of http response is %d instead of 200", response.StatusCode)
		return
	}

	return
}

func (c *Client) dump(response *http.Response) *DumpData {
	dumpData := new(DumpData)

	if response == nil {
		return dumpData
	}

	request := response.Request

	// Dump request
	dumpRequest, err := httputil.DumpRequest(request, false)
	dumpData.Request = string(dumpRequest)
	dumpData.DumpRequestError = err.Error()

	// Dump request body
	requestBody, err := request.GetBody()
	defer func() {
		if requestBody != nil {
			requestBody.Close()
		}
	}()
	if err != nil {
		dumpData.DumpRequestBodyError = err.Error()
	} else {
		requestBodyBytes, err := io.ReadAll(requestBody)
		dumpData.RequestBody = string(requestBodyBytes)
		dumpData.DumpRequestBodyError = err.Error()
	}

	// Dump response
	dumpResponse, err := httputil.DumpResponse(response, false)
	dumpData.Response = string(dumpResponse)
	dumpData.DumpResponseError = err.Error()

	// Dump response body
	responseBody := response.Body
	defer func() {
		if responseBody != nil {
			responseBody.Close()
		}
	}()
	responseBodyBytes, err := io.ReadAll(responseBody)
	dumpData.ResponseBody = string(responseBodyBytes)
	dumpData.DumpResponseBodyError = err.Error()

	return dumpData
}
