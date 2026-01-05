package main

import (
	"fmt"
	"time"
)

type Ride struct {
	ID       string    `json:"id"`
	Time     time.Time `json:"time"`
	Distance float64
	Shared   bool
}

func (r Ride) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("empty ID")
	}

	if r.Time.IsZero() {
		return fmt.Errorf("missing time")
	}

	if r.Distance <= 0 {
		return fmt.Errorf("%f - bad distance", r.Distance)
	}

	return nil
}
