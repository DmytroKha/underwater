package main

import (
	"context"
	"fmt"
	"github.com/DmytroKha/underwater/config"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/DmytroKha/underwater/internal/infra/database"
	"github.com/DmytroKha/underwater/internal/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

var previousTransparency = make(map[int64]int64)
var allFish []string

func main() {

	exitCode := 0
	_, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	var conf = config.GetConfiguration()

	err := database.Migrate(conf)
	if err != nil {
		log.Fatalf("Unable to apply migrations: %q\n", err)
	}

	generateFishSpecies()

	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}

	//Repository
	sensorRepository := database.NewSensorRepository(sess)

	//Service
	sensorService := app.NewSensorService(sensorRepository)

	sensors, err := sensorService.FindAll()

	if err != nil {
		log.Print(err)
		return
	}

	for _, sensor := range sensors {
		go generateSensorData(sensor)
	}
	select {}

}

func generateSensorData(sensor models.Sensor) {
	for {
		reading := generateFakeReading(sensor)
		log.Println(reading)
		//saveReadingToDatabase(reading)
		time.Sleep(time.Duration(sensor.DataOutputRate) * time.Second)
	}
}

func generateFakeReading(sensor models.Sensor) models.Reading {
	return models.Reading{
		SensorID:     sensor.ID,
		Temperature:  generateFakeTemperature(sensor.Z),
		Transparency: generateFakeTransparency(sensor.GroupID),
		FishSpecies:  generateFakeFishSpecies(),
		Timestamp:    time.Now(),
	}
}

func generateFakeTemperature(depth float64) float64 {
	return 10.0 + (depth / 10) + rand.Float64()*5
}

func generateFakeTransparency(groupID int64) int64 {
	maxChange := 5

	previous, exists := previousTransparency[groupID]
	if !exists {
		previous = int64(rand.Intn(100))
	}

	newTransparency := previous + int64(rand.Intn(2*maxChange+1)) - int64(maxChange)

	if newTransparency < 0 {
		newTransparency = 0
	} else if newTransparency > 100 {
		newTransparency = 100
	}

	previousTransparency[groupID] = newTransparency

	return newTransparency
}

func generateFakeFishSpecies() []models.FishSpecies {

	var fishSpecies []models.FishSpecies
	numFishSpecies := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	fishCount := make(map[string]int64)

	for i := 0; i < numFishSpecies; i++ {
		randomIndex := rand.Intn(len(allFish))
		randomFish := allFish[randomIndex]
		fishCount[randomFish]++
	}

	for f, c := range fishCount {
		fish := models.FishSpecies{
			Name:  f,
			Count: c,
		}
		fishSpecies = append(fishSpecies, fish)
	}

	return fishSpecies
}

func generateFishSpecies() {

	url := "https://oceana.org/ocean-fishes/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.tb-grid-column h2").Each(func(index int, item *goquery.Selection) {
		fishName := item.Text()
		allFish = append(allFish, fishName)
	})

}
