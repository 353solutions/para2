package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL string
	c       http.Client
}

func NewClient(baseURL string) Client {
	return Client{baseURL: baseURL}
}

func (c *Client) Health() error {
	url, err := url.JoinPath(c.baseURL, "/health")
	if err != nil {
		return err
	}

	resp, err := c.c.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: bad status - %s", url, resp.Status)
	}

	return nil
}
