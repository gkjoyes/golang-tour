package main

import "fmt"

type animal struct {
	name   string
	origin string
}

type bird struct {
	animal
	speed  float32
	canFly bool
}

func main() {
	b := bird{
		animal: animal{
			name:   "Emu",
			origin: "Australia",
		},
		speed:  48,
		canFly: false,
	}
	fmt.Printf("%+v\n", b)
}
