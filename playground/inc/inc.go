package main

import "fmt"

func main() {
	n := 41
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
	inc(&n)
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
	fmt.Println("n:", n)

	s := []int{10, 20, 30, 40}
	// s1 := s[:3]
	s1 := s[:3:3] // Also sets capacity
	fmt.Println("s1 len:", len(s1), "cap:", cap(s1))

	/*
		// fmt.Println(s1[3]) // panic
		s4 := s1[:4] // You can slice up to cap
		fmt.Println(s4)
	*/
	s2 := []int{100} //, 200}
	// concat two slices
	out := append(s1, s2...)
	fmt.Println("out:", out)
	fmt.Println("s:", s)

	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("arr: %v %T\n", arr, arr)

	// slice to array
	sa := []int{10, 20, 30, 40}
	var arr2 [4]int
	copy(arr2[:], sa)
	fmt.Println(arr2)
}

func inc(n *int) {
	*n++
	fmt.Printf("inc: n=%d, addr=%p\n", *n, n)
}

/* Go passes by value, but:
- map, channel are pointers
- slice contains a pointer

Prefer value semantics. You must use pointer if:
- Changing struct fields
- Lock fields (sync.Mutex) - protobuf
- Unmarshaling

Prefer to stay in same semantics
*/
