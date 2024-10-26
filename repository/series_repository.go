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

func (repository *SeriesInfoRepository) Add(series models.Series) (int64, error) {
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

func (repository *SeriesInfoRepository) GetAll() ([]models.Series, error) {
	rows, err := repository.db.Query("SELECT * FROM series ORDER BY title LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("error querying series from the database")
	}
	defer rows.Close()

	var series []models.Series
	for rows.Next() {
		var s models.Series

		err := rows.Scan(&s.ID,
			&s.Title,
			&s.Link,
			&s.IMDB,
			&s.StartYear,
			&s.EndYear,
			&s.Poster,
			&s.Country)
		if err != nil {
			return nil, fmt.Errorf("error scanning series data")
		}

		series = append(series, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over series rows")
	}

	return series, nil
}

func (repository *SeriesInfoRepository) GetById(id int) (models.Series, error) {
	row := repository.db.QueryRow("SELECT * FROM series WHERE id = :id",
		sql.Named("id", id))
	var series models.Series

	err := row.Scan(&series.ID,
		&series.Title,
		&series.Link,
		&series.IMDB,
		&series.StartYear,
		&series.EndYear,
		&series.Poster,
		&series.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			return series, fmt.Errorf("series not found")
		} else {
			return series, fmt.Errorf("error retrieving series from database")
		}
	}

	return series, nil
}
