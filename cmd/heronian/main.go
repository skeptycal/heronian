package main

import (
	"fmt"

	"github.com/skeptycal/heronian"
)

func main() {

	t := heronian.New(3, 4, 5)

	fmt.Printf("t#: %#v\n", t)
	fmt.Printf("t: %s\n", t)
	fmt.Printf("t.Sides(): %v\n", t.Sides())
	fmt.Printf("Perimeter: %v\n", t.Perimeter())
	fmt.Printf("Area: %v\n", t.Area())
	fmt.Printf("SemiPerimeter: %v\n", t.SemiPerimeter())
	fmt.Printf("Heron: %v\n", t.Heron())
	fmt.Printf("IsHero: %v\n", t.IsHero())

	s := heronian.NewShape(3, 4, 5)
	fmt.Printf("s#: %#v\n", s)
	fmt.Printf("s: %s\n", s)

}
