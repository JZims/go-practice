package main

import (
	"fmt"
)

func main() {
	chihuahua := Dog{"Chihuahua", 4}

	fmt.Println(chihuahua)
	fmt.Printf("%+v\n", chihuahua)

}

type Dog struct {
	Breed  string
	Weight int
}
