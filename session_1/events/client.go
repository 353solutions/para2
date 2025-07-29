package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	/*
		e := events[0]
		data, err := json.Marshal(e)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		r := bytes.NewReader(data)
	*/
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	go func() {
		defer w.Close()
		enc := json.NewEncoder(w)
		for _, e := range events {
			if err := enc.Encode(e); err != nil {
				slog.Error("encode", "error", err)
				return
			}
		}
	}()

	const url = "http://localhost:8080/events"
	// BUG: No timeout
	// resp, err := http.Post(url, "application/json", r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, r)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer s3cr3t")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR:", resp.Status)
		return
	}

	io.Copy(os.Stdout, resp.Body)

}

var events = []Event{
	{asTime("2025-05-21T14:31:49Z"), "elliot", "read", "file:///etc/passwd"},
	{asTime("2025-05-21T14:42:32Z"), "elliot", "read", "file:///etc/shadow"},
	{asTime("2025-05-21T14:43:07Z"), "elliot", "read", "file:///root/.ssh/config"},
}

type Event struct {
	Time   time.Time `json:"time"`
	Login  string    `json:"login"`
	Action string    `json:"action"`
	URI    string    `json:"uri"`
}

func asTime(s string) time.Time {
	t, _ := time.Parse(s, time.RFC3339)
	return t
}
