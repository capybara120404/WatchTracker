package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara120404/watch-tracker/models"
)

func main() {
	url := "https://s2.fanserialstv.net"
	var yesterdaySeriesInfo, todaySeriesInfo []models.SeriesInfo
	var err error
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go func() {
		for {
			todaySeriesInfo, err = fetchSeriesData(url)
			if err != nil {
				log.Println(err)
			}
			if !reflect.DeepEqual(todaySeriesInfo, yesterdaySeriesInfo) {
				yesterdaySeriesInfo = todaySeriesInfo

				detailsWG := sync.WaitGroup{}
				for i := range todaySeriesInfo {
					detailsWG.Add(1)
					go func(series *models.SeriesInfo) {
						defer detailsWG.Done()
						fetchSeriesDetails(series, url)
					}(&todaySeriesInfo[i])
				}
				detailsWG.Wait()
			}

			for i := 0; i < 10; i++ {
				fmt.Println(todaySeriesInfo[i])
			}

			time.Sleep(24 * time.Hour)
		}
	}()
}

func fetchSeriesData(url string) ([]models.SeriesInfo, error) {
	var seriesData []models.SeriesInfo

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

		seriesData = append(seriesData, models.SeriesInfo{
			Title:     title,
			Link:      url + link,
			IMDB:      imdb,
			StartYear: startYear,
			EndYear:   endYear,
		})
	})

	return seriesData, nil
}

func fetchSeriesDetails(series *models.SeriesInfo, url string) {
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
	poster ,_:= doc.Find(".field-poster img").Attr("src")
	series.Poster = url+poster
	series.Country = doc.Find(".info-list li:contains('Страна:') .field-text").Text()
}
