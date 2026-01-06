package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestAPI_Health(t *testing.T) {
	ctx := t.Context()
	c, err := testcontainers.Run(
		ctx, "",
		testcontainers.WithExposedPorts("8080/tcp"),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("8080/tcp"),
		),
		testcontainers.WithDockerfile(testcontainers.FromDockerfile{
			Context:    ".",
			Dockerfile: "Dockerfile",
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer testcontainers.CleanupContainer(t, c)

	host, err := c.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := c.MappedPort(ctx, "8080/tcp")
	if err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("http://%s:%s/health", host, port.Port())
	t.Logf("url: %v", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("bad status: %s", resp.Status)
	}
}
