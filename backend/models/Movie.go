package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate string    `json:"release_date"`
	RunTime     int       `json:"runtime"`
	Image       string    `json:"image"`
	MPAARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Genre       *Genre    `json:genre,omitempty`
}

type Genre struct {
	ID        int       `json:"id"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
