package database

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const ReadingTableName = "sensor_readings"

type reading struct {
	ID           int64     `db:"id,omitempty"`
	SensorID     int64     `db:"sensor_id"`
	Timestamp    time.Time `db:"timestamp"`
	Temperature  float64   `db:"temperature"`
	Transparency int64     `db:"transparency"`
}

type ReadingRepository interface {
	Save(reading domain.Reading) (domain.Reading, error)
	GetAverageTemperatureBySensor(sensorID int64, from time.Time, till time.Time) (float64, error)
	GetAverageTemperatureForGroup(sensorsIds []int64) (float64, error)
	GetAverageTransparencyForGroup(sensorsIds []int64) (float64, error)
}

type readingRepository struct {
	coll db.Collection
	sess db.Session
}

func NewReadingRepository(dbSession db.Session) ReadingRepository {
	return readingRepository{
		coll: dbSession.Collection(ReadingTableName),
		sess: dbSession,
	}
}

func (r readingRepository) Save(reading domain.Reading) (domain.Reading, error) {
	re := r.mapDomainToModel(reading)

	err := r.coll.InsertReturning(&re)
	if err != nil {
		return domain.Reading{}, err
	}

	return r.mapModelToDomain(re), nil
}

func (r readingRepository) GetAverageTemperatureBySensor(sensorID int64, fromDate time.Time, tillDate time.Time) (float64, error) {
	var averageTemperature float64

	var result struct {
		Temperature float64 `db:"temperature"`
	}

	query := r.sess.SQL().Select(db.Raw("AVG(temperature) AS temperature")).From("sensor_readings").
		Where(db.Cond{
			"sensor_id =":  sensorID,
			"timestamp >=": fromDate,
			"timestamp <=": tillDate,
		})

	err := query.One(&result)
	if err != nil {
		return 0, err
	}

	averageTemperature = result.Temperature

	return averageTemperature, nil
}

func (r readingRepository) GetAverageTemperatureForGroup(sensorsIDs []int64) (float64, error) {
	var averageTemperature float64

	var result struct {
		Temperature float64 `db:"temperature"`
	}

	query := r.sess.SQL().Select(db.Raw("AVG(temperature) AS temperature")).From("sensor_readings").
		Where(db.Cond{"sensor_id IN": sensorsIDs})

	err := query.One(&result)
	if err != nil {
		return 0, err
	}

	averageTemperature = result.Temperature

	return averageTemperature, nil
}

func (r readingRepository) GetAverageTransparencyForGroup(sensorsIDs []int64) (float64, error) {
	var averageTransparency float64

	var result struct {
		Transparency float64 `db:"transparency"`
	}

	query := r.sess.SQL().Select(db.Raw("AVG(transparency) AS transparency")).From("sensor_readings").
		Where(db.Cond{"sensor_id IN": sensorsIDs})

	err := query.One(&result)
	if err != nil {
		return 0, err
	}

	averageTransparency = result.Transparency

	return averageTransparency, nil
}

func (r readingRepository) mapDomainToModel(re domain.Reading) reading {
	return reading{
		ID:           re.ID,
		SensorID:     re.SensorID,
		Timestamp:    re.Timestamp,
		Temperature:  re.Temperature,
		Transparency: re.Transparency,
	}
}

func (r readingRepository) mapModelToDomain(re reading) domain.Reading {
	return domain.Reading{
		ID:           re.ID,
		SensorID:     re.SensorID,
		Timestamp:    re.Timestamp,
		Temperature:  re.Temperature,
		Transparency: re.Transparency,
	}
}

func (r readingRepository) mapModelToDomainCollection(readings []reading) []domain.Reading {
	result := make([]domain.Reading, len(readings))

	for i := range readings {
		result[i] = r.mapModelToDomain(readings[i])
	}

	return result
}
