package models

import "fmt"

type SeriesInfo struct {
	ID            int
	Title         string
	Link          string
	IMDB          string
	StartYear     string
	EndYear       string
	Poster        string
	Country       string
	NumberOfViews int
}

func (s SeriesInfo) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Link %s, IMDB: %s, StartYear: %s, EndYear: %s, Poster: %s, Country: %s, NumberOfViews: %d",
		s.ID,
		s.Title,
		s.Link,
		s.IMDB,
		s.StartYear,
		s.EndYear,
		s.Poster,
		s.Country,
		s.NumberOfViews)
}
