package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
	"log"
	"math/rand"
	"sync"
	"time"
)

type ReadingService interface {
	GenerateSensorData(sensor domain.Sensor)
	GetAverageTemperatureBySensor(sensorID int64, from int64, till int64) (float64, error)
	GetAverageTemperatureForGroup(sensorsIDs []int64) (float64, error)
	GetAverageTransparencyForGroup(sensorsIDs []int64) (float64, error)
	GetRegionMinTemperatureBySensors(sensorsIDs []int64) (float64, error)
	GetRegionMaxTemperatureBySensors(sensorsIDs []int64) (float64, error)
	FindBySensorsIDs(sensorsIDs []int64, from int64, till int64) ([]domain.Reading, error)
}

type readingService struct {
	readingRepo        database.ReadingRepository
	sensorService      SensorService
	fishSpeciesService FishSpeciesService
}

var previousTransparency = make(map[int64]int64)

var transparencyMutex sync.Mutex

func NewReadingService(r database.ReadingRepository, ss SensorService, fs FishSpeciesService) ReadingService {
	return readingService{
		readingRepo:        r,
		sensorService:      ss,
		fishSpeciesService: fs,
	}
}

func (s readingService) GetAverageTemperatureBySensor(sensorID int64, from int64, till int64) (float64, error) {
	fromDate := time.Unix(from, 0)
	tillDate := time.Unix(till, 0)
	return s.readingRepo.GetAverageTemperatureBySensor(sensorID, fromDate, tillDate)
}

func (s readingService) GetAverageTemperatureForGroup(sensorsIDs []int64) (float64, error) {
	return s.readingRepo.GetAverageTemperatureForGroup(sensorsIDs)
}

func (s readingService) GetAverageTransparencyForGroup(sensorsIDs []int64) (float64, error) {
	return s.readingRepo.GetAverageTransparencyForGroup(sensorsIDs)
}

func (s readingService) GetRegionMinTemperatureBySensors(sensorsIDs []int64) (float64, error) {
	return s.readingRepo.GetRegionMinTemperatureBySensors(sensorsIDs)
}

func (s readingService) GetRegionMaxTemperatureBySensors(sensorsIDs []int64) (float64, error) {
	return s.readingRepo.GetRegionMaxTemperatureBySensors(sensorsIDs)
}

func (s readingService) FindBySensorsIDs(sensorsIDs []int64, from int64, till int64) ([]domain.Reading, error) {
	fromDate := time.Unix(from, 0)
	tillDate := time.Unix(till, 0)
	return s.readingRepo.FindBySensorsIDs(sensorsIDs, fromDate, tillDate)
}

func (s readingService) GenerateSensorData(sensor domain.Sensor) {
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
		Timestamp:    time.Now(),
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
