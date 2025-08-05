package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_HealthError(t *testing.T) {
	var c Client
	c.c.Transport = errTripper{}
	err := c.Health(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}

}

// RoundTrip implements http.RoundTripper
func (errTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("no network")
}

type errTripper struct{}
