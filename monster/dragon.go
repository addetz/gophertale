package monster

import "fmt"

type Dragon struct {
	Name string
	Message string
}

func (d Dragon) DoBadStuff() {
	d.roar()
}

func (d Dragon) roar() {
	fmt.Printf("Raaaah! I'm the %s!\n", d.Name)
	fmt.Println(d.Message)
}
