package main

import (
	"log"
)

func main() {
	a, err := App.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
