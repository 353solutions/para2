package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for _, l := range logs {
		if err := enc.Encode(l); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	}

	dec := json.NewDecoder(&buf)
loop: // label
	for {
		var l Log
		err := dec.Decode(&l)
		/*
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
		*/
		switch {
		case errors.Is(err, io.EOF):
			break loop
		case err != nil:
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Println(l)
	}
}

type Log struct {
	Host    string    `json:"host"`
	Time    time.Time `json:"time"`
	Request string    `json:"request"`
	Status  int       `json:"status"`
	Bytes   int       `json:"bytes"`
}

var logs = []Log{
	{"199.72.81.55", asTime("2023-07-01T00:00:01Z"), "GET /history/apollo/ HTTP/1.1", 200, 6245},
	{"unicomp6.unicomp.net", asTime("2023-07-01T00:00:06Z"), "GET /shuttle/countdown/ HTTP/1.1", 200, 3985},
	{"199.120.110.21", asTime("2023-07-01T00:00:09Z"), "GET /shuttle/missions/sts-73/mission-sts-73.html HTTP/1.1", 200, 4085},
	{"burger.letters.com", asTime("2023-07-01T00:00:11Z"), "GET /shuttle/countdown/liftoff.html HTTP/1.1", 304, 0},
	{"199.120.110.21", asTime("2023-07-01T00:00:11Z"), "GET /shuttle/missions/sts-73/sts-73-patch-small.gif HTTP/1.1", 200, 4179},
	{"burger.letters.com", asTime("2023-07-01T00:00:12Z"), "GET /images/NASA-logosmall.gif HTTP/1.1", 304, 0},
	{"burger.letters.com", asTime("2023-07-01T00:00:12Z"), "GET /shuttle/countdown/video/livevideo.gif HTTP/1.1", 200, 0},
	{"205.212.115.106", asTime("2023-07-01T00:00:12Z"), "GET /shuttle/countdown/countdown.html HTTP/1.1", 200, 3985},
	{"d104.aa.net", asTime("2023-07-01T00:00:13Z"), "GET /shuttle/countdown/ HTTP/1.1", 200, 3985},
	{"129.94.144.152", asTime("2023-07-01T00:00:13Z"), "GET / HTTP/1.1", 200, 7074},
}

func asTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}
