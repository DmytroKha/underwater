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
	FishSpecies  []byte    `db:"fish_species"`
}

type ReadingRepository interface {
	Save(reading domain.Reading) (domain.Reading, error)
	//FindAll() ([]models.Sensor, error)
}

type readingRepository struct {
	coll db.Collection
}

func NewReadingRepository(dbSession db.Session) ReadingRepository {
	return readingRepository{
		coll: (dbSession).Collection(ReadingTableName),
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

func (r readingRepository) mapDomainToModel(re domain.Reading) reading {
	//fishSpeciesJSON, err := json.Marshal(re.FishSpecies)
	//if err != nil {
	//	fishSpeciesJSON = nil
	//}
	return reading{
		ID:           re.ID,
		SensorID:     re.SensorID,
		Timestamp:    re.Timestamp,
		Temperature:  re.Temperature,
		Transparency: re.Transparency,
		//FishSpecies:  fishSpeciesJSON,
	}
}

func (r readingRepository) mapModelToDomain(re reading) domain.Reading {
	//var fishSpecies []domain.FishSpecies
	//err := json.Unmarshal(re.FishSpecies, fishSpecies)
	//if err != nil {
	//	fishSpecies = []domain.FishSpecies{}
	//}
	return domain.Reading{
		ID:           re.ID,
		SensorID:     re.SensorID,
		Timestamp:    re.Timestamp,
		Temperature:  re.Temperature,
		Transparency: re.Transparency,
		//FishSpecies:  fishSpecies,
	}
}

func (r readingRepository) mapModelToDomainCollection(readings []reading) []domain.Reading {
	result := make([]domain.Reading, len(readings))

	for i := range readings {
		result[i] = r.mapModelToDomain(readings[i])
	}

	return result
}
