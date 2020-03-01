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
	req.Header.Set("Cookie", `ll="118282"; bid=8EVK7RUIBzQ; _vwo_uuid_v2=D190AA8A5C73CB4421FDAC74C94840733|b96ce9d21112ab7981ba2393f150ef8b; douban-fav-remind=1; __utmv=30149280.6697; _ga=GA1.2.862603563.1555553241; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1567305915%2C%22https%3A%2F%2Fwww.google.com%2F%22%5D; __utma=30149280.862603563.1555553241.1567003981.1567305916.9; __utmz=30149280.1567305916.9.7.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); _pk_id.100001.8cb4=75cc8b425932499a.1557154780.4.1567305922.1564290022.; push_doumail_num=0;dbcl2="66975456:cXxEDrJitwU"; ck=jwbA; push_noty_num=0; ap_v=0,6.0`)
	req.Header.Set("Referer", "https://www.douban.com/group/")

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

