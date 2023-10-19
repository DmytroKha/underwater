package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"sync"
	"time"
)

type ReadingService interface {
	StartSensorDataGeneration()
	generateSensorData(sensor domain.Sensor)
	GetAverageTemperatureBySensor(sensorID int64, from int64, till int64) (float64, error)
}

type readingService struct {
	readingRepo        database.ReadingRepository
	sensorService      SensorService
	fishSpeciesService FishSpeciesService
}

var previousTransparency = make(map[int64]int64)
var allFish []string
var transparencyMutex sync.Mutex

func NewReadingService(r database.ReadingRepository, ss SensorService, fs FishSpeciesService) ReadingService {
	return readingService{
		readingRepo:        r,
		sensorService:      ss,
		fishSpeciesService: fs,
	}
}

func (s readingService) StartSensorDataGeneration() {
	generateFishSpecies()
	sensors, err := s.sensorService.FindAll()
	if err != nil {
		log.Print(err)
		return
	}

	for _, sensor := range sensors {
		go s.generateSensorData(sensor)
	}
	select {}
}

func (s readingService) GetAverageTemperatureBySensor(sensorID int64, from int64, till int64) (float64, error) {
	fromDate := time.Unix(from, 0)
	tillDate := time.Unix(till, 0)
	return s.readingRepo.GetAverageTemperatureBySensor(sensorID, fromDate, tillDate)
}

func (s readingService) generateSensorData(sensor domain.Sensor) {
	for {
		reading := generateFakeReading(sensor)
		transparencyMutex.Lock()
		re, err := s.readingRepo.Save(reading)
		transparencyMutex.Unlock()
		if err != nil {
			log.Print(err)
		}
		reID := re.ID
		transparencyMutex.Lock()
		err = s.fishSpeciesService.Save(reID)
		transparencyMutex.Unlock()
		if err != nil {
			log.Print(err)
		}

		time.Sleep(time.Duration(sensor.DataOutputRate) * time.Second)
	}
}

func generateFakeReading(sensor domain.Sensor) domain.Reading {
	return domain.Reading{
		SensorID:     sensor.ID,
		Temperature:  generateFakeTemperature(sensor.Z),
		Transparency: generateFakeTransparency(sensor.GroupID),
		//FishSpecies:  generateFakeFishSpecies(),
		Timestamp: time.Now(),
	}
}

func generateFakeTemperature(depth float64) float64 {
	return 10.0 + (depth / 10) + rand.Float64()*5
}

func generateFakeTransparency(groupID int64) int64 {
	maxChange := 5

	transparencyMutex.Lock()
	previous, exists := previousTransparency[groupID]
	transparencyMutex.Unlock()
	if !exists {
		previous = int64(rand.Intn(100))
	}

	newTransparency := previous + int64(rand.Intn(2*maxChange+1)) - int64(maxChange)

	if newTransparency < 0 {
		newTransparency = 0
	} else if newTransparency > 100 {
		newTransparency = 100
	}

	transparencyMutex.Lock()
	previousTransparency[groupID] = newTransparency
	transparencyMutex.Unlock()

	return newTransparency
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
