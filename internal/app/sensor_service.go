package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
)

type SensorService interface {
	FindAll() ([]domain.Sensor, error)
}

type sensorService struct {
	sensorRepo database.SensorRepository
}

func NewSensorService(r database.SensorRepository) SensorService {
	return sensorService{
		sensorRepo: r,
	}
}

func (s sensorService) FindAll() ([]domain.Sensor, error) {
	return s.sensorRepo.FindAll()
}
