package main

import (
	"encoding/json"
	"fmt"
	"go-refactor/ch1"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("ch1/invoices.json")
	if err != nil {
		panic(err)
	}
	var invoice ch1.Invoice
	err = json.Unmarshal(bytes, &invoice)
	if err != nil {
		panic(err)
	}
	plays := make(map[string]ch1.Play, 0)
	plays["hamlet"] = ch1.Play{
		Name: "Hamlet",
		Type: "tragedy",
	}

	plays["as-like"] = ch1.Play{
		Name: "As You Like It",
		Type: "comedy",
	}

	plays["othello"] = ch1.Play{
		Name: "Othello",
		Type: "tragedy",
	}

	statement := ch1.V1Handler.Statement(invoice, plays)
	fmt.Print(statement)
}
