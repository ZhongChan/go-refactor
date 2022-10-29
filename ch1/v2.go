package ch1

import (
	"fmt"
	"math"
)

var V2Handler = V2{}

type V2 struct {
}

func (V2) Statement(invoices Invoice, plays map[string]Play) string {
	totalAmount := 0.0
	volumeCredits := 0.0
	result := fmt.Sprintf("Statement for %s\n", invoices.Customer)

	for _, perf := range invoices.Performances {
		if _, ok := plays[perf.PlayID]; !ok {
			continue
		}
		play := plays[perf.PlayID]

		thisAmount := 0.0
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
		default:
			panic(fmt.Sprintf(`unknown type:%s`, play.Type))
		}

		volumeCredits += math.Max(perf.Audience-30, 0)
		if "comedy" == play.Type {
			volumeCredits += math.Floor(perf.Audience / 5)
		}
		result += fmt.Sprintf(" %s:$%.2f (%0.f seats)\n", play.Name, thisAmount/100, perf.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("You earned %0.f credits\n", volumeCredits)
	return result
}
