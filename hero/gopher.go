package hero

import "fmt"

type Gopher struct {
	Name string
}

func (g Gopher) Enter() {
	fmt.Printf("Out of nowhere, the famous hero, %s, appears.\n", g.Name)
	fmt.Println("\"I will help you all escape this monster!\", the hero confidently says.")
}

func (g Gopher) Save() {
	fmt.Printf("%s took the little codebase up in the clouds, away from all monsters.\n", g.Name)
}
