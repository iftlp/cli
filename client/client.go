package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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
		Status   string
		Message  string
		Response *http.Response
	}
)

func (c *Client) do(method, endpoint string, params map[string]string, data []byte) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)

	req, err := http.NewRequest(method, baseURL, bytes.NewBuffer(data))
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

func (c *Client) SendWavFile(path string) (resp Response, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("audioFile", "audio.wav")
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	err = writer.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/recognition/wavFile", c.host), &buf)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := c.httpClient.Do(req)

	bufBody := new(strings.Builder)

	io.Copy(bufBody, res.Body)

	resp.Status = res.Status
	resp.Message = bufBody.String()
	resp.Response = res

	return
}

func (c *Client) SendMusicByte(datas []byte) (resp Response, err error) {
	jsonData, err := json.Marshal(datas)
	if err != nil {
		panic(err)
	}

	log.Println(jsonData)

	res, err := c.do(http.MethodPost, "recognition/rawData", nil, jsonData)
	if err != nil {
		return
	}

	buf := new(strings.Builder)

	io.Copy(buf, res.Body)

	resp.Status = res.Status
	resp.Message = buf.String()
	resp.Response = res

	return
}

func CreateClient(host string) *Client {
	client := &http.Client{
		Timeout: time.Second * 100,
	}
	return &Client{
		host:       host,
		httpClient: client,
	}
}
