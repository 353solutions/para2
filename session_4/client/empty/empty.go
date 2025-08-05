package main

import "fmt"

func main() {
	var a any // Go < 1.18 interface{}

	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a:", a)

	s := a.(string) // type assertion
	fmt.Println("s:", s)

	// i := a.(int) // panic
	if i, ok := a.(int); ok {
		fmt.Println("i:", i)
	} else {
		fmt.Println("Not an int")
	}

	switch v := a.(type) { // type switch
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
		fmt.Println(v + "!")
	default:
		fmt.Printf("unknown type: %T\n", a)
	}

	/*
		switch {
		case v := msg.GetPaypal(); v != nil:
			//...
		}
	*/

	// See reflect package, in general avoid it
}

/*
message Event {
	Login login = 1;
	Logout logout = 2;
	Access access = 3;
}
*/

/*
Rule of thumb: Don't use any (interface{})
Exceptions:
- Printing
- Marshaling (serialization) (json.Marshal ...)
*/
