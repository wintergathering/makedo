package models

type Bathroom struct {
	ID     uint   `json:"id"`
	Place  string `json:"place"`
	Rating uint   `json:"rating"`
	Review string `json:"review"`
}
