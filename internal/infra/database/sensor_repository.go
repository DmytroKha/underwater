package database

import (
	"github.com/DmytroKha/underwater/internal/models"
	"github.com/upper/db/v4"
	"log"
)

const SensorsTableName = "sensors"

type sensor struct {
	ID             int64   `db:"id,omitempty"`
	Codename       string  `db:"codename"`
	GroupID        int64   `db:"group_id"`
	X              float64 `db:"x"`
	Y              float64 `db:"y"`
	Z              float64 `db:"z"`
	DataOutputRate int64   `db:"data_output_rate"`
}

type SensorRepository interface {
	FindAll() ([]models.Sensor, error)
}

type sensorRepository struct {
	coll db.Collection
}

func NewSensorRepository(dbSession db.Session) SensorRepository {
	return sensorRepository{
		coll: (dbSession).Collection(SensorsTableName),
	}
}

func (r sensorRepository) FindAll() ([]models.Sensor, error) {
	var sensors []sensor

	err := r.coll.Find().All(&sensors)

	if err != nil {
		log.Print(err)
		return []models.Sensor{}, err
	}

	return r.mapModelToDomainCollection(sensors), nil
}

func (r sensorRepository) mapModelToDomain(s sensor) models.Sensor {
	return models.Sensor{
		ID:             s.ID,
		Codename:       s.Codename,
		GroupID:        s.GroupID,
		X:              s.X,
		Y:              s.Y,
		Z:              s.Z,
		DataOutputRate: s.DataOutputRate,
	}
}

func (r sensorRepository) mapModelToDomainCollection(sensors []sensor) []models.Sensor {
	result := make([]models.Sensor, len(sensors))

	for i := range sensors {
		result[i] = r.mapModelToDomain(sensors[i])
	}

	return result
}
