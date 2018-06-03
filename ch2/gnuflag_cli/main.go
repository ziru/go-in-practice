package main

import (
	"fmt"

	flag "github.com/juju/gnuflag"
)

var name = flag.String("name", "World", "A name to say hello to")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	flag.Parse(false)
	if spanish == true {
		fmt.Printf("Hola %s! Args: %v\n", *name, flag.Args())
	} else {
		fmt.Printf("Hello %s! Args: %v\n", *name, flag.Args())
	}
}
