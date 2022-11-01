package ch1

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID   string  `json:"playID"`
	Audience float64 `json:"audience"`
}
