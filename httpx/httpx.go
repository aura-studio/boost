package httpx

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
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
	var request = func() (data []byte, err error) {
		resp, err := doFunc(c.Client)
		defer func() {
			if resp != nil {
				resp.Body.Close()
			}
		}()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("status code of http response is %d instead of 200", resp.StatusCode)
		}
		return io.ReadAll(resp.Body)
	}

	var times int64
	for times < c.Config.Retry {
		times++
		data, err = request()
		if err != nil {
			continue
		}
		return data, nil
	}

	return nil, err
}
