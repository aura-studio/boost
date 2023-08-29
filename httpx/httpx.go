package httpx

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DoFunc func(*http.Client) (*http.Response, error)

type ClientConfig struct {
	Proxy   string
	Retry   int64
	Timeout time.Duration
}

type Client struct {
	Config ClientConfig
	*http.Client
}

var DefaultClientConfig = ClientConfig{
	Proxy:   "",
	Retry:   1,
	Timeout: 0,
}

func Default() *Client {
	return NewClient(DefaultClientConfig)
}

func NewClient(config ClientConfig) *Client {
	c := &Client{
		Config: config,
		Client: &http.Client{},
	}

	if c.Config.Proxy != "" {
		proxyURL, err := url.Parse(c.Config.Proxy)
		if err != nil {
			log.Panicln(err)
		}

		c.Client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
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
