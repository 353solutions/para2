package client

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Health(ctx context.Context) error {
	url := fmt.Sprintf("%s/health", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%q - %s", url, resp.Status)
	}

	return nil
}

type Client struct {
	baseURL string
	c       http.Client
}
