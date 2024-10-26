package models

import "fmt"

type User struct {
	ID       int
	UserName string
	Email    string
	Age      int
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, UserName: %s, Email: %s, Age: %d",
		u.ID,
		u.UserName,
		u.Email,
		u.Age)
}

type Series struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Link      string `json:"link"`
	IMDB      string `json:"imdb,omitempty"`
	StartYear string `json:"start_year,omitempty"`
	EndYear   string `json:"end_year,omitempty"`
	Poster    string `json:"poster,omitempty"`
	Country   string `json:"country,omitempty"`
}

func (s Series) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Link %s, IMDB: %s, StartYear: %s, EndYear: %s, Poster: %s, Country: %s",
		s.ID,
		s.Title,
		s.Link,
		s.IMDB,
		s.StartYear,
		s.EndYear,
		s.Poster,
		s.Country)
}

type UserSeries struct {
	UserID   int
	SeriesID int
	Views    int
}

func (u UserSeries) String() string {
	return fmt.Sprintf("UserID: %d, SeriesID: %d, Views: %d", u.UserID, u.SeriesID, u.Views)
}
