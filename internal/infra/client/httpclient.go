package client

import (
	"time"
	"net/http"
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
