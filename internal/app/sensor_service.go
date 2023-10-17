package app

import (
	"github.com/DmytroKha/underwater/internal/infra/database"
	"github.com/DmytroKha/underwater/internal/models"
)

type SensorService interface {
	FindAll() ([]models.Sensor, error)
}

type sensorService struct {
	sensorRepo database.SensorRepository
}

func NewSensorService(r database.SensorRepository) SensorService {
	return sensorService{
		sensorRepo: r,
	}
}

func (s sensorService) FindAll() ([]models.Sensor, error) {
	return s.sensorRepo.FindAll()
}
