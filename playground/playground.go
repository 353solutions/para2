package main

import (
	"fmt"
	"time"
)

/*
Go
type Reader interface {
	Read([]byte) (int, error)
}

Python
type Reader interface {
	Read(int) ([]byte, error)
}
*/

func main() {

	/*
		a, b := 1.1, 2.2
		fmt.Println(a + b)
	*/

	/*
		s := "שלום"
		fmt.Println(len(s))   // bytes
		for _, c := range s { // runes
			fmt.Printf("%c\n", c)
		}
	*/

	/*
		var s []int
		c := 0
		for i := range 10_000 {
			s = append(s, i)
			if cap(s) > c {
				fmt.Println(c, "->", cap(s))
				c = cap(s)
			}
		}
	*/
	/*
		m := map[string]int{
			"A": 47,
			"B": 36,
		}
		v := m["C"]
		fmt.Println(v)

		v, ok := m["C"]
		fmt.Println(v, ok)
		return
	*/
	/*
		//csvFile := "c:\to\new\reports\2025.csv"
		// `a` is a "raw string", \ is just a \
		csvFile := `c:\to\new\reports\2025.csv`
		fmt.Println(csvFile)
		return
	*/

	ch := make(chan int)
	go func() {
		v := <-ch
		fmt.Println("v:", v)
		time.Sleep(time.Second)
		fmt.Println("done")
	}()
	ch <- 1
	fmt.Println("send")
	time.Sleep(2 * time.Second)

}
