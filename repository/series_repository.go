package repository

import (
	"database/sql"
	"fmt"

	"github.com/capybara120404/watch-tracker/database"
	"github.com/capybara120404/watch-tracker/models"
)

type SeriesInfoRepository struct {
	db *sql.DB
}

func NewSeriesInfoRepository(connecter *database.Connecter) *SeriesInfoRepository {
	return &SeriesInfoRepository{
		db: connecter.DB,
	}
}

func (repository *SeriesInfoRepository) Add(series models.SeriesInfo) (int64, error) {
	res, err := repository.db.Exec("INSERT INTO series (title, link, imdb, start_year, end_year, poster, country) VALUES (:title, :link, :imdb, :start_year, :end_year, :poster, :country)",
		sql.Named("title", series.Title),
		sql.Named("link", series.Link),
		sql.Named("imdb", series.IMDB),
		sql.Named("start_year", series.StartYear),
		sql.Named("end_year", series.EndYear),
		sql.Named("poster", series.Poster),
		sql.Named("country", series.Country))
	if err != nil {
		return 0, fmt.Errorf("error inserting data into the database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error retrieving last insert Id")
	}

	return id, nil
}
