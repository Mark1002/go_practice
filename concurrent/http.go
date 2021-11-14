package concurrent

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExecuteHttp() {
	result := crawl("hello-world")
	fmt.Println(result)
}

func crawl(path string) string {
	resp, err := http.Get("https://gobyexample.com/" + path)
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
	title := doc.Find("h2").Text()
	return title
}
