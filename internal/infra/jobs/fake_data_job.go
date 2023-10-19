package jobs

import (
	"github.com/DmytroKha/underwater/internal/app"
	"log"
)

//var AllFish []string

func StartSensorDataGeneration(readingService *app.ReadingService, sensorService *app.SensorService) {
	//generateFishSpecies()
	sensors, err := (*sensorService).FindAll()
	if err != nil {
		log.Print(err)
		return
	}

	for _, sensor := range sensors {
		go (*readingService).GenerateSensorData(sensor)
	}
	select {}
}

//func generateFishSpecies() {
//
//	url := "https://oceana.org/ocean-fishes/"
//
//	doc, err := goquery.NewDocument(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	doc.Find("div.tb-grid-column h2").Each(func(index int, item *goquery.Selection) {
//		fishName := item.Text()
//		AllFish = append(AllFish, fishName)
//	})
//
//}
