package main

import (
	"log"

	"github.com/goedgar/pkg/edgar"
	"github.com/goedgar/pkg/platform/http"
)

const GOLDMAN_SACHS_CIK = "886982"

func main() {
	client := edgar.New(
		*http.New(),
	)

	// fs, err := client.FilingsForCIK(GOLDMAN_SACHS_CIK)
	// if err != nil {
	// 	panic(err)
	// }

	// docs, err := client.DocsForFiling(GOLDMAN_SACHS_CIK, fs[0].Name)
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := client.FileForDoc(GOLDMAN_SACHS_CIK, fs[0].Name, docs[3].Name)
	// if err != nil {
	// 	panic(err)
	// }

	items, err := client.DailyIndex("2019")
	if err != nil {
		panic(err)
	}

	log.Println(items)
}
