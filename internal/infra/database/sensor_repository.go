package database

import (
	"github.com/DmytroKha/underwater/internal/domain"
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
	//Save(reading domain.Reading) error
	FindAll() ([]domain.Sensor, error)
}

type sensorRepository struct {
	coll db.Collection
}

func NewSensorRepository(dbSession db.Session) SensorRepository {
	return sensorRepository{
		coll: (dbSession).Collection(SensorsTableName),
	}
}

//func (r sensorRepository) Save(reading models.Reading) error {
//	u := r.mapDomainToModel(reading)
//
//	err := r.coll.InsertReturning(&u)
//	if err != nil {
//		return domain.User{}, err
//	}
//
//	return r.mapModelToDomain(u), nil
//}

func (r sensorRepository) FindAll() ([]domain.Sensor, error) {
	var sensors []sensor

	err := r.coll.Find().All(&sensors)

	if err != nil {
		log.Print(err)
		return []domain.Sensor{}, err
	}

	return r.mapModelToDomainCollection(sensors), nil
}

func (r sensorRepository) mapDomainToModel(s domain.Sensor) sensor {
	return sensor{
		ID:             s.ID,
		Codename:       s.Codename,
		GroupID:        s.GroupID,
		X:              s.X,
		Y:              s.Y,
		Z:              s.Z,
		DataOutputRate: s.DataOutputRate,
	}
}

func (r sensorRepository) mapModelToDomain(s sensor) domain.Sensor {
	return domain.Sensor{
		ID:             s.ID,
		Codename:       s.Codename,
		GroupID:        s.GroupID,
		X:              s.X,
		Y:              s.Y,
		Z:              s.Z,
		DataOutputRate: s.DataOutputRate,
	}
}

func (r sensorRepository) mapModelToDomainCollection(sensors []sensor) []domain.Sensor {
	result := make([]domain.Sensor, len(sensors))

	for i := range sensors {
		result[i] = r.mapModelToDomain(sensors[i])
	}

	return result
}
