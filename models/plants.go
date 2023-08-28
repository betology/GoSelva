package models

import "time"

type Plant struct {
	ID          string    `json:"id"`
	Genus       string    `json:"genus"`
	Species     string    `json:"species"`
	PublishedAt time.Time `json:"publishedAt"`
	Picture     string    `json:"imageURL"`
	Family      string    `json:"family"`
	Location    string    `json:"location"`
	Names       []string  `json:"names"`
	AcquiredAt  time.Time `json:"acquiredAt"`
	Notes       []string  `json:"notes"`
}
