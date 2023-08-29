package httpx

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type DoFunc func(*http.Client) (*http.Response, error)

type ClientConfig struct {
	Proxy string
	Retry int64
}

type Client struct {
	ClientConfig
	client *http.Client
}

var DefaultClientConfig = ClientConfig{
	Proxy: "",
	Retry: 1,
}

func Default() *Client {
	return NewClient(DefaultClientConfig)
}

func NewClient(config ClientConfig) *Client {
	c := &Client{
		ClientConfig: config,
	}

	if c.Proxy == "" {
		c.client = &http.Client{}
	} else {
		proxyURL, err := url.Parse(c.Proxy)
		if err != nil {
			log.Panicln(err)
		}

		c.client = &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}}
	}

	return c
}

func Request(doFunc DoFunc) ([]byte, error) {
	return Default().Request(doFunc)
}

func (c *Client) Request(doFunc DoFunc) (data []byte, err error) {
	var request = func() (data []byte, err error) {
		resp, err := doFunc(c.client)
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
	for times < c.Retry {
		times++
		data, err = request()
		if err != nil {
			continue
		}
		return data, nil
	}

	return nil, err
}
