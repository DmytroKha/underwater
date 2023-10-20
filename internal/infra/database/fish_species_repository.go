package database

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/upper/db/v4"
)

const FishSpeciesTableName = "fish_species"

type fishSpecies struct {
	ID        int64  `db:"id,omitempty"`
	ReadingID int64  `db:"reading_id"`
	Name      string `db:"name"`
	Count     int64  `db:"count"`
}

type FishSpeciesRepository interface {
	Save(reading domain.FishSpecies) error
	GetFishesForGroup(readingsIDs []int64) ([]domain.FishSpecies, error)
	GetTopFishesForGroup(readingsIDs []int64, fishCount int64) ([]domain.FishSpecies, error)
}

type fishSpeciesRepository struct {
	coll db.Collection
	sess db.Session
}

func NewFishSpeciesRepository(dbSession db.Session) FishSpeciesRepository {
	return fishSpeciesRepository{
		coll: (dbSession).Collection(FishSpeciesTableName),
		sess: dbSession,
	}
}

func (r fishSpeciesRepository) Save(fishSpecies domain.FishSpecies) error {
	fs := r.mapDomainToModel(fishSpecies)

	err := r.coll.InsertReturning(&fs)
	if err != nil {
		return err
	}

	return nil
}

func (r fishSpeciesRepository) GetFishesForGroup(readingsIDs []int64) ([]domain.FishSpecies, error) {
	var fs []fishSpecies

	query := r.sess.SQL().Select(db.Raw("name, SUM(count) AS count")).From("fish_species").
		Where(db.Cond{"reading_id IN": readingsIDs}).
		GroupBy("name").
		OrderBy("name")
	//OrderBy("count DESC")

	err := query.All(&fs)

	if err != nil {
		return []domain.FishSpecies{}, err
	}

	return r.mapModelToDomainCollection(fs), nil
}

func (r fishSpeciesRepository) GetTopFishesForGroup(readingsIDs []int64, fishCount int64) ([]domain.FishSpecies, error) {
	var fs []fishSpecies

	//dbCond := db.Cond{}
	//dbCond["reading_id IN"] = readingsIDs
	//if fromDate != tillDate {
	//	dbCond["timestamp >="] = fromDate
	//	dbCond["timestamp <="] = tillDate
	//
	//}

	query := r.sess.SQL().Select(db.Raw("name, SUM(count) AS count")).From("fish_species").
		Where(db.Cond{"reading_id IN": readingsIDs}).
		GroupBy("name").
		OrderBy("count DESC").
		Limit(int(fishCount))

	err := query.All(&fs)

	if err != nil {
		return []domain.FishSpecies{}, err
	}

	return r.mapModelToDomainCollection(fs), nil
}

func (r fishSpeciesRepository) mapDomainToModel(fs domain.FishSpecies) fishSpecies {
	return fishSpecies{
		ID:        fs.ID,
		ReadingID: fs.ReadingID,
		Name:      fs.Name,
		Count:     fs.Count,
	}
}

func (r fishSpeciesRepository) mapModelToDomain(fs fishSpecies) domain.FishSpecies {
	return domain.FishSpecies{
		ID:        fs.ID,
		ReadingID: fs.ReadingID,
		Name:      fs.Name,
		Count:     fs.Count,
	}
}

func (r fishSpeciesRepository) mapModelToDomainCollection(fishSpecies []fishSpecies) []domain.FishSpecies {
	result := make([]domain.FishSpecies, len(fishSpecies))

	for i := range fishSpecies {
		result[i] = r.mapModelToDomain(fishSpecies[i])
	}

	return result
}
