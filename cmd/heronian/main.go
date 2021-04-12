package main

import (
	"fmt"

	"github.com/skeptycal/cli"
	"github.com/skeptycal/heronian"
)

func main() {

	out := cli.NewStdout(nil)

	out.Println("Hero Triangle Example")
	out.Br()

	t := heronian.New(1, 1, 1)

	fmt.Printf("t#: %#v\n", t)
	fmt.Printf("t: %s\n", t)
	fmt.Printf("t.Sides(): 		%30v\n", t.Sides())
	fmt.Printf("Perimeter: 		%30v\n", t.Perimeter())
	fmt.Printf("SemiPerimeter: 	%30v\n", t.SemiPerimeter())
	fmt.Printf("Area: 			%30v\n", t.Area())
	fmt.Printf("Heron: 			%30v\n", t.HeronArea())
	fmt.Printf("IsHero: 		%30v\n", t.IsHero())

	s := heronian.NewShape(3, 4, 5)
	fmt.Printf("s#: %#v\n", s)
	fmt.Printf("s: %s\n", s)

}
