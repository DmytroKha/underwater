package database

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/upper/db/v4"
)

const GroupTableName = "sensor_groups"

type group struct {
	ID   int64  `db:"id,omitempty"`
	Name string `db:"name"`
}

type GroupRepository interface {
	FindByName(groupName string) (domain.SensorGroup, error)
}

type groupRepository struct {
	coll db.Collection
}

func NewGroupRepository(dbSession db.Session) GroupRepository {
	return groupRepository{
		coll: (dbSession).Collection(GroupTableName),
	}
}

func (r groupRepository) FindByName(groupName string) (domain.SensorGroup, error) {
	var g group

	err := r.coll.Find(db.Cond{"name": groupName}).One(&g)
	if err != nil {
		return domain.SensorGroup{}, err
	}

	return r.mapModelToDomain(g), nil
}

func (r groupRepository) mapDomainToModel(g domain.SensorGroup) group {
	return group{
		ID:   g.ID,
		Name: g.Name,
	}
}

func (r groupRepository) mapModelToDomain(s group) domain.SensorGroup {
	return domain.SensorGroup{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (r groupRepository) mapModelToDomainCollection(groups []group) []domain.SensorGroup {
	result := make([]domain.SensorGroup, len(groups))

	for i := range groups {
		result[i] = r.mapModelToDomain(groups[i])
	}

	return result
}
