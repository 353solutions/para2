package main

import "fmt"

func main() {
	i, err := NewItem(200, 300)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(i)
	fmt.Println(NewItem(200, 4000))

	// i3 := Item{-1, 20_000}

	i.Move(10, 30)
	fmt.Println("i (move):", i)

	p1 := Player{
		Name: "Parzival",
	}
	// Embedding lifts fields & methods to embedding type
	fmt.Println("p1.X:", p1.X)
	fmt.Println("p1.Item.X:", p1.Item.X)
	p1.Move(207, 33)
	// C++: p1->Move(207, 33);
	fmt.Println("p1 (move):", p1)

	ms := []Mover{
		i,
		&p1, // must match receiver semantics
	}
	moveAll(ms, 37, 123)
	for _, m := range ms {
		fmt.Println("m:", m)
	}

	fmt.Println("jade:", Jade)
	fmt.Printf("jade: %d\n", Jade)
	fmt.Println("Key(18):", Key(18))
}

/* What fmt.Print* does
func Print(v any) {
	if s, ok := v.(fmt.Stringer) {
		print(s.String())
		return
	}

	switch v.(type) {
	case int:
		// ...
	}
}
*/

// String implements fmt.Stringer
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("Key(%d)", k)
}

// Go's "enum"
type Key uint8

const (
	Copper Key = iota + 1 // 1 << iota // bitmask
	Jade
	Crystal
)

/* Interface
- Say what we need, not what we provide
- Interfaces are small (stdlib average < 2 methods), my rule is up to 4
- Accept interfaces, return types

Uses:
- Polymorphism
	- Mocking
- How stdlib looks at your types
*/

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type Mover interface {
	Move(int, int)
}

// value -> pointer ✓
// pointer -> value ✗

type Player struct {
	Name string
	Item // Player embeds Item

	// X *int // masks Item.X
}

// "i" is called "the receiver"
// Use pointer receiver if you change fields
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// 2D game

/* New/Factory functions
func NewItem(x, y int) Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) *Item
func NewItem(x, y int) (*Item, error)
*/

/* If you want set/getter
Make x, y unexported
Getter: X()
Setter: SetX(x int)
*/

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of range for %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// Go does "escape analysis" and allocates i on the heap
	// go build -gcflags=-m
	return &i, nil
}

const (
	maxX = 400
	maxY = 600
)

type Item struct {
	X int
	Y int
}
