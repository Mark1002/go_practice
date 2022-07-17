package concurrent

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type HttpCrawlerDemo struct{}

func (dmo HttpCrawlerDemo) Execute() {
	targets := crawlLinks()
	fmt.Println("target count:", len(targets))
	results := make([]string, 0)
	messages := make(chan string)
	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		go crawlTitle(target, messages)
		go func() {
			defer wg.Done()
			results = append(results, <-messages)
		}()
	}
	wg.Wait()
	fmt.Println("results count:", len(results))
}

func crawlByGOQuey(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func crawlTitle(path string, messages chan<- string) {
	url := "https://gobyexample.com/"
	doc := crawlByGOQuey(url + path)
	title := doc.Find("h2").Text()
	messages <- title
}

func crawlLinks() []string {
	url := "https://gobyexample.com/"
	links := make([]string, 0)
	doc := crawlByGOQuey(url)
	doc.Find("li > a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			links = append(links, link)
		}
	})
	return links
}
