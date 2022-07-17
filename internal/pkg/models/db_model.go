package models

// Collector collector object structure
type Content struct {
	UserName     string `json:"Username"`
	SortKey      string `json:"sortKey"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	Organization string `json:"org"`
}
