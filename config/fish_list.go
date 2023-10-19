package config

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

var AllFish []string

func GenerateFishSpecies() {

	url := "https://oceana.org/ocean-fishes/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.tb-grid-column h2").Each(func(index int, item *goquery.Selection) {
		fishName := item.Text()
		AllFish = append(AllFish, fishName)
	})

}
