package main

import (
	"encoding/json"
	"fmt"
	"go-refactor/ch1"
	"io/ioutil"
)

func main() {
	invoicesBytes, err := ioutil.ReadFile("ch1/invoices.json")
	if err != nil {
		panic(err)
	}
	var invoice ch1.Invoice
	err = json.Unmarshal(invoicesBytes, &invoice)
	if err != nil {
		panic(err)
	}

	playsBytes, err := ioutil.ReadFile("ch1/plays.json")
	if err != nil {
		panic(err)
	}
	plays := make(map[string]ch1.Play, 0)
	err = json.Unmarshal(playsBytes, &plays)
	if err != nil {
		panic(err)
	}

	fmt.Print(ch1.V1Handler.Statement(invoice, plays))
	fmt.Print(ch1.V2Handler.Statement(invoice, plays))
}
