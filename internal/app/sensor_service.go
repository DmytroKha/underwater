package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
)

type SensorService interface {
	FindAll() ([]domain.Sensor, error)
	FindByName(codeName string) (domain.Sensor, error)
	FindByGroupID(groupID int64) ([]domain.Sensor, error)
	GetAverageTemperatureBySensor(codeName string, from int64, till int64) (float64, error)
}

type sensorService struct {
	sensorRepo     database.SensorRepository
	readingService *ReadingService
}

func NewSensorService(r database.SensorRepository, rs *ReadingService) SensorService {
	return sensorService{
		sensorRepo:     r,
		readingService: rs,
	}
}

func (s sensorService) FindAll() ([]domain.Sensor, error) {
	return s.sensorRepo.FindAll()
}

func (s sensorService) FindByGroupID(groupID int64) ([]domain.Sensor, error) {
	return s.sensorRepo.FindByGroupID(groupID)
}

func (s sensorService) FindByName(codeName string) (domain.Sensor, error) {
	return s.sensorRepo.FindByName(codeName)
}

func (s sensorService) GetAverageTemperatureBySensor(codeName string, from int64, till int64) (float64, error) {
	sensor, err := s.FindByName(codeName)
	if err != nil {
		return 0, err
	}
	return (*s.readingService).GetAverageTemperatureBySensor(sensor.ID, from, till)
}
