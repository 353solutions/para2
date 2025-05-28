package main

import "time"

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
