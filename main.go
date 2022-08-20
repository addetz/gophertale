package main

import (
	"fmt"

	"github.com/addetz/gophertale/hero"
	"github.com/addetz/gophertale/monster"
)

type Baddie interface {
	DoBadStuff()
}

func main() {
	name := "New Requirements Dragon"
	message := "This is my home now! You all must go live in the clouds before sundown."
	fmt.Println("Once upon a time, there was a little codebase who lived in the village server.")
	fmt.Println("One day, unexpectedly, a ghastly monster darkened the village sky.")

	d := monster.Dragon{
		Name:    name,
		Message: message,
	}
	w := monster.Witch{
		Name:    "High latency witch",
		Message: "I curse you with high latency for all your web traffic!",
	}
	baddies := []Baddie{d, w}
	for _, b := range baddies {
		b.DoBadStuff()
	}

	fmt.Println("The little codebase trembled.")
	fmt.Println("Is there a hero out there that can help us??")

	g := hero.Gopher{
		Name: "Super Gopherlina",
	}
	g.Enter()
	g.Save()

	fmt.Println("And they all lived happily ever after. The end.")
}
