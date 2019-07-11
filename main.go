package main

import (
	"log"

	"github.com/goedgar/edgar"
)

const GOLDMAN_SACHS_CIK = "886982"

func main() {
	client := edgar.New()

	fs, err := client.FilingsForCIK(GOLDMAN_SACHS_CIK)
	if err != nil {
		panic(err)
	}

	docs, err := client.FilingDocs(GOLDMAN_SACHS_CIK, fs[0].Name)
	if err != nil {
		panic(err)
	}

	log.Println(docs[0])

}
