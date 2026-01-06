package main_test

import (
	"fmt"
	unter "unter"
)

func ExampleRidePrice() {
	price := unter.RidePrice(5.0, true)
	fmt.Println(price)

	// Output:
	// 900
}
