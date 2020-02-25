package douban

import (
	"fmt"
	"io"
	"net/http"
)

const (
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"
	host             = "https://douban.com"
)

// http client
type Client struct {
	//some payload for http request
	*http.Client
}

func NewClient() *Client {
	return &Client {
		Client: &http.Client{
			Transport:     http.DefaultTransport,
		},
	}
}

func (c *Client) NewRequest(method, url string) (*http.Request, error){

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Host", host)

	return req, nil
}

func (c *Client) Do(r *http.Request) (io.ReadCloser, error){
	response, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, fmt.Errorf("error with request, %s %d %s", Url, response.StatusCode, response.Status)
	}
	return response.Body, nil
}

