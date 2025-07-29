package main

import "fmt"

func main() {
	if _, err := OpenFile("/dev/null"); err != nil {
		// Can use errors.Is, errors.As to inspect err
		fmt.Println("ERROR:", err)
	}
	fmt.Println("OK")

	var i *LoginMessage
	fmt.Println(i.Kind())
}

// It's ok that data field in interface is nil
func (*LoginMessage) Kind() string {
	return "Item"
}

type LoginMessage struct{}

/* interface definition  src/runtime/runtime2.go
type iface struct {
	tab  *itab  		// pointer to all implement type methods
	data unsafe.Pointer // pointer to actual value
}

var err *OSError
	data -> nil
	tab -> &OSError

Interface is nil iff both data are nil
*/

func OpenFile(path string) (*File, error) {
	// var err *OSError // BUG: err as error is not nil
	var err error

	// TODO:
	return nil, err
}

type File struct{}

// Errors implements error
func (o *OSError) Error() string {
	return o.Path
}

type OSError struct {
	Path string
}
