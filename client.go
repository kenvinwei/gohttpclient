package gohttpclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewClient(t int) *Client {
	return client(t)
}

func client(t int) *Client {
	return &Client{
		TimeOut: time.Duration(t),
	}
}

func (c *Client) SetHeader(headers map[string]string) *Client {
	if len(headers) == 0 {
		return c
	}

	c.H.Header = http.Header{}

	for k, v := range headers {
		c.H.Header.Add(k, v)
	}

	c.hasHeader = true

	return c
}

func (c *Client) Get(url string) []byte {
	var (
		req    *http.Request
		resp   *http.Response
		err    error
		result []byte
	)

	cli := http.Client{
		Timeout: c.TimeOut * time.Second,
	}

	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		log.Fatalf("http.NewRequest err:%v\n", err)
	}

	//add header
	if c.hasHeader {
		req.Header = c.H.Header
	}

	if resp, err = cli.Do(req); err != nil {
		log.Fatalf("cli.Do err:%v\n", err)
	}

	defer resp.Body.Close()

	//log.Printf("get resp request:%v\n", resp.Request)

	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalf("ioutil.ReadAll err:%v\n", err)
	}

	return result
}

func (c *Client) ParseForm(form FormValue) string {
	if len(form) == 0 {
		return ""
	}

	c.D = make(url.Values)

	for k, v := range form {
		if len(v) == 0 {
			continue
		}

		if len(v) == 1 {
			c.D.Set(k, v[0])
			continue
		}

		for _, vv := range v {
			c.D.Add(k, vv)
		}
	}

	return c.D.Encode()
}

func (c *Client) Post(url string, form FormValue) []byte {
	var (
		req    *http.Request
		resp   *http.Response
		err    error
		result []byte
	)

	cli := http.Client{}

	if c.TimeOut > 0 {
		cli.Timeout = c.TimeOut * time.Second
	}

	PostByteReader := strings.NewReader("")
	if len(form) > 0 {
		PostDataStr := c.ParseForm(form)

		PostByteReader = strings.NewReader(PostDataStr)
	}

	if req, err = http.NewRequest(http.MethodPost, url, PostByteReader); err != nil {
		log.Fatalf("http.NewRequest err:%v\n", err)
	}

	//add header
	if c.hasHeader {
		req.Header = c.H.Header
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if resp, err = cli.Do(req); err != nil {
		log.Fatalf("cli.Do err:%+v\n", err)
	}

	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalf("ioutil.ReadAll err:%v\n", err)
	}

	return result
}
