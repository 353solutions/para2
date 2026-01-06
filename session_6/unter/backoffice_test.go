package main

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/require"
)

func TestRidePrice(t *testing.T) {
	price := RidePrice(5, false)
	expected := 1000
	require.Equal(t, expected, price)
	/*
		if price != expected {
			t.Fatalf("expected: %#v, got %#v", expected, price)
		}
	*/
}

type priceTestCase struct {
	Distance float64
	Shared   bool
	Expected int
	Name     string
}

// TODO: return iter.Seq[priceTestCase]
func loadPriceTestCases(t *testing.T) []priceTestCase {
	// t.Helper()
	file, err := os.Open("testdata/price_cases.yml")
	require.NoError(t, err, "file")
	defer file.Close()

	var cases []priceTestCase
	err = yaml.NewDecoder(file).Decode(&cases)
	require.NoError(t, err, "decode")
	return cases
}

func TestRidePrice_Table(t *testing.T) {
	for _, tc := range loadPriceTestCases(t) {
		name := tc.Name
		if name == "" {
			name = fmt.Sprintf("%v-%v", tc.Distance, tc.Shared)
		}
		t.Run(name, func(t *testing.T) {
			price := RidePrice(tc.Distance, tc.Shared)
			require.Equal(t, tc.Expected, price)
		})
	}
}

func FuzzRidePrice(f *testing.F) {
	f.Add(0.0, false)
	f.Fuzz(func(t *testing.T, distance float64, shared bool) {
		distance = math.Abs(distance)
		price := RidePrice(distance, shared)
		if price <= 0 {
			t.Fatalf("%v:%v -> %d - bad price", distance, shared, price)
		}
	})
}

// $ go test -v -fuzz . -fuzztime 10s
