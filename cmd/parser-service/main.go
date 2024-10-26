package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara120404/watch-tracker/database"
	"github.com/capybara120404/watch-tracker/models"
	"github.com/capybara120404/watch-tracker/repository"
)

const url string = "https://s2.fanserialstv.net"

func main() {
	connecter, err := database.OpenOrCreate("watch_tracker.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer connecter.Close()

	var series []models.Series
	series, err = fetchSeriesData(url)
	if err != nil {
		log.Println(err)
		return
	}

	wg := sync.WaitGroup{}
	for i := range series {
		wg.Add(1)
		go func(series *models.Series) {
			defer wg.Done()
			fetchSeriesDetails(series, url)
		}(&series[i])
	}
	wg.Wait()

	repository := repository.NewSeriesInfoRepository(connecter)
	for _, s := range series {
		repository.Add(s)
	}
}

func fetchSeriesData(url string) ([]models.Series, error) {
	var seriesData []models.Series

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("page request error: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error when loading HTML: %v", err)
	}

	doc.Find("li.literal__item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		link, exists := s.Find("a").Attr("href")
		if !exists {
			link = ""
		}
		imdb, _ := s.Attr("data-imdb")
		startYear, _ := s.Attr("data-start-year")
		endYear, _ := s.Attr("data-end-year")

		seriesData = append(seriesData, models.Series{
			Title:     title,
			Link:      url + link,
			IMDB:      imdb,
			StartYear: startYear,
			EndYear:   endYear,
		})
	})

	return seriesData, nil
}

func fetchSeriesDetails(series *models.Series, url string) {
	resp, err := http.Get(series.Link)
	if err != nil {
		log.Printf("error fetching series details for %s: %v", series.Title, err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("error parsing series details for %s: %v", series.Title, err)
		return
	}
	poster, _ := doc.Find(".field-poster img").Attr("src")
	series.Poster = url + poster
	series.Country = doc.Find(".info-list li:contains('Страна:') .field-text").Text()
}
