package httpclient

import (
	"fmt"
	"time"
)

type HttpClient struct {
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) GetData(url string, amount int) (string, error) {
	time.Sleep(time.Millisecond * 200)
	msg := fmt.Sprintf("search for %d at %s failed", url, amount)
	return msg, nil
}
