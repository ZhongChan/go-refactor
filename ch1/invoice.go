package ch1

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []struct {
		PlayID   string  `json:"playID"`
		Audience float64 `json:"audience"`
	} `json:"performances"`
}

