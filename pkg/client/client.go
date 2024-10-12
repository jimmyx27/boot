package client

import (
	"errors"
	"fmt"
)

type HttpClient interface {
	GetData(url string, amount int) (string, error)
}

type Logger interface {
	Debug(msg string)
	Error(err error)
}

type Client struct {
	logger     Logger
	httpClient HttpClient
}

func NewClient(logger Logger, httpClient HttpClient) (*Client, error) {
	if logger == nil {
		return nil, errors.New("NewClient called but logger is nil")
	}
	if httpClient == nil {
		return nil, errors.New("NewClient called but httpClient is nil")
	}

	return &Client{
		logger:     logger,
		httpClient: httpClient,
	}, nil
}

func (c *Client) GetData(url string, amount int) (string, error) {
	c.logger.Debug(fmt.Sprintf("GetData called with %s and %d", url, amount))
	return c.httpClient.GetData(url, amount)
}
