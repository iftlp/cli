package client

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type (
	Client struct {
		host       string
		httpClient *http.Client
	}
)

type (
	Response struct {
		Status  string
		Message string
	}
)

func (c *Client) do(method, endpoint string, params map[string]string, body []byte) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

func (c *Client) SendMusicByte(datas []byte) (resp Response, err error) {
	res, err := c.do(http.MethodPost, "recognition/load", nil, datas)
	if err != nil {
		return
	}

	resp.Status = res.Status
	resp.Message = "true"

	return
}

func CreateClient(host string) *Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return &Client{
		host:       host,
		httpClient: client,
	}
}
