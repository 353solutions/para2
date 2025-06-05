package store

import (
	"fmt"
	"os"
)

func ExampleStore() {
	file, err := os.CreateTemp("", "*.bbolt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	file.Close()

	db, err := Open(file.Name())
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer db.Close()

	s := Store{db}
	i := Item{
		SKU:   "sku17",
		Name:  "Gopher Plushy",
		Price: 375,
	}
	fmt.Println(s.Update(i))

	i2, err := s.GetItem(i.SKU)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(i2)

	i.Price = 350
	fmt.Println(s.Update(i))
	i2, err = s.GetItem(i.SKU)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(i2)

	// Output:
	// <nil>
	// {sku17 Gopher Plushy 375}
	// <nil>
	// {sku17 Gopher Plushy 350}
}
