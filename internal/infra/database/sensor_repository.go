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
	FindAll() ([]domain.Sensor, error)
	FindByName(codeName string) (domain.Sensor, error)
	FindByGroupID(groupID int64) ([]domain.Sensor, error)
	FindByCoordinates(xMin, yMin, zMin, xMax, yMax, zMax float64) ([]domain.Sensor, error)
}

type sensorRepository struct {
	coll db.Collection
}

func NewSensorRepository(dbSession db.Session) SensorRepository {
	return sensorRepository{
		coll: (dbSession).Collection(SensorsTableName),
	}
}

func (r sensorRepository) FindAll() ([]domain.Sensor, error) {
	var sensors []sensor

	err := r.coll.Find().All(&sensors)

	if err != nil {
		log.Print(err)
		return []domain.Sensor{}, err
	}

	return r.mapModelToDomainCollection(sensors), nil
}

func (r sensorRepository) FindByName(codeName string) (domain.Sensor, error) {
	var s sensor

	err := r.coll.Find(db.Cond{"codename": codeName}).One(&s)
	if err != nil {
		return domain.Sensor{}, err
	}

	return r.mapModelToDomain(s), nil
}

func (r sensorRepository) FindByGroupID(groupID int64) ([]domain.Sensor, error) {
	var s []sensor

	err := r.coll.Find(db.Cond{"group_id": groupID}).All(&s)
	if err != nil {
		return []domain.Sensor{}, err
	}

	return r.mapModelToDomainCollection(s), nil
}

func (r sensorRepository) FindByCoordinates(xMin, yMin, zMin, xMax, yMax, zMax float64) ([]domain.Sensor, error) {
	var s []sensor

	dbCond := db.Cond{}
	dbCond["x >="] = xMin
	dbCond["y >="] = yMin
	dbCond["z >="] = zMin
	dbCond["x <="] = xMax
	dbCond["y <="] = yMax
	dbCond["z <="] = zMax

	err := r.coll.Find(dbCond).All(&s)
	if err != nil {
		return []domain.Sensor{}, err
	}

	return r.mapModelToDomainCollection(s), nil
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
