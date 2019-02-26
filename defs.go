package gohttpclient

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	H         Header
	hasHeader bool
	Method    string
	Url       string
	TimeOut   time.Duration
	D         url.Values
}

type FormValue map[string][]string

type Header struct {
	http.Header
}
