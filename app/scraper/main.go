package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	model "github.com/Mahamadou828/AOAC/business/data/v1/models/university"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/sys/database"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/gocolly/colly"
)

var availableEnv = map[string]string{
	"local":      "local",
	"testing":    "testing",
	"staging":    "staging",
	"production": "production",
}

/*
* @todo insert a timeout for the scraping, if the timeout is exceeded we should close the channel and move on
 */

func main() {
	//Get and check the environment that run the scrapper.
	env, ok := os.LookupEnv("ENV")
	if !ok {
		log.Fatalf("forget to specify environment variable: env")
	}
	if _, ok := availableEnv[env]; !ok {
		log.Fatalf("unvalid environment: %v", env)
	}

	//Initialize the db connection
	client, err := aws.New(aws.Config{
		ServiceName: "university-scraper",
		Environment: env,
	})
	if err != nil {
		log.Fatalf("can't initialize a aws session: %v", err)
	}
	db := database.Open(client, env)
	ctx := context.Background()

	//Create a new collector that will be share and use throughout the application
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.SetRequestTimeout(30 * time.Second)
	cu, ce := make(chan model.University), make(chan error)

	//
	var wg sync.WaitGroup
	wg.Add(274)

	go func() {
		wg.Wait()
		close(cu)
	}()

	go func() {
		select {
		case err := <-ce:
			fmt.Println(err)
		}
	}()

	for i := 1; i < 275; i++ {
		go scrap(i, &wg, c.Clone(), cu, ce)
	}

	var us []model.University

	for u := range cu {
		log.Println("saving university of", u.Name)
		//completing the university struct with additional fields
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		u.ID = validate.GenerateID()

		us = append(us, u)

		if err := model.Create(ctx, db, u); err != nil {
			log.Println("failing to save university", u.Name, "with error", err)
		}
	}

	file, _ := json.MarshalIndent(us, "", "")
	if err := os.WriteFile("response.json", file, 0644); err != nil {
		log.Fatalf("failed to backup universities in response.json: %v", err)
	}
}

func scrap(i int, wg *sync.WaitGroup, c *colly.Collector, cu chan<- model.University, ce chan<- error) {
	url := fmt.Sprintf("https://www.hotcoursesabroad.com/study/international/schools-colleges-university/list.html?sortby=ALL&pageNo=%d", i)

	c.OnHTML(
		".pr_rslt",
		func(e *colly.HTMLElement) {
			var u model.University
			u.Name = strings.ToLower(e.ChildText(".sr_nam > h2"))
			log.Println("scraping university", u.Name)
			u.Country = strings.ToLower(e.ChildText(".sr_nam > span.grey"))
			u.Rating = extractRating(e.ChildAttr(".pr_hd > .sr_rvw > .rvw_ratg > i:not(.fa)", "class"))
			u.DetailsURL = e.Attr("href")

			if len(u.Name) <= 0 || len(u.Country) <= 0 {
				return
			}

			scrapUniversityDetail(u.DetailsURL, c.Clone(), func(faculties []string, desc string) {
				u.Faculties = faculties
				u.Description = desc
				log.Println("sending university to save", u.Name)
				cu <- u
			})
		})

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.OnError(func(response *colly.Response, err error) {
		ce <- fmt.Errorf("error visiting %s: %v", url, err)
		wg.Done()
	})

	if err := c.Visit(url); err != nil {
		ce <- fmt.Errorf("error visiting %s: %v", url, err)
		wg.Done()
	}
}

func extractRating(rtg string) int {
	switch strings.ToLower(rtg) {
	case "rating5":
		return 5
	case "rating4":
		return 4
	case "rating3":
		return 3
	case "rating2":
		return 2
	case "rating1":
		return 1
	default:
		return 0
	}
}

func scrapUniversityDetail(url string, c *colly.Collector, callback func(faculties []string, desc string)) {
	log.Println("scraping details:", url)
	var desc string
	var faculties []string
	c.OnHTML(
		"body",
		func(e *colly.HTMLElement) {
			e.ForEach(".chub_cont_col > #nav-section-6 > ul > li", func(i int, e *colly.HTMLElement) {
				faculties = append(faculties, e.Text)
			})

			desc = e.ChildText("#nav-section-1 > p:nth-child(3)")
		})

	c.OnScraped(func(r *colly.Response) {
		callback(faculties, desc)
	})
	c.Visit(url)
}
