package monster

import "fmt"

type Witch struct {
	Name    string
	Message string
}

func (w Witch) DoBadStuff() {
	w.curse()
}

func (w Witch) curse() {
	fmt.Printf("Heckle heckle! I'm the %s!\n", w.Name)
	fmt.Println(w.Message)
}
