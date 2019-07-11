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

	content, err := client.FilingDocContent(GOLDMAN_SACHS_CIK, fs[0].Name, docs[0].Name)
	if err != nil {
		panic(err)
	}
	log.Println(content)

	// for _, doc := range docs {
	// 	if doc.Type == edgar.TextType {
	// 		content, err := client.FilingDocContent(GOLDMAN_SACHS_CIK, fs[0].Name, doc.Name)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Println(content)
	// 	}
	// }

}
