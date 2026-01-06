package main

import "math"

// RidePrice returns ride price in ¢
func RidePrice(distance float64, shared bool) int {
	price := 250 // initial fare
	price += int(math.Ceil(distance)) * 150

	if shared {
		price = int(float64(price) * 0.9)
	}

	return price
}
