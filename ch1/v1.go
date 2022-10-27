package ch1

import "fmt"

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

type Play struct {
	Hamlet struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"hamlet"`
	AsLike struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"as-like"`
	Othello struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"othello"`
}

func Statement(invoices Invoice, plays Play) {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoices.Customer)
	fmt.Println(totalAmount, volumeCredits, result)
}
