package main

import (
	"log"

	"github.com/goedgar/edgar"
)

func main() {
	client := edgar.New()

	fs, err := client.FilingsForCIK("886982")
	if err != nil {
		panic(err)
	}

	log.Println(fs[0])
}
