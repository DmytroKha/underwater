package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
)

type GroupService interface {
	GetAverageTemperatureForGroup(groupName string) (float64, error)
	GetAverageTransparencyForGroup(groupName string) (float64, error)
	GetFishesForGroup(groupName string) ([]domain.FishSpecies, error)
	GetTopFishesForGroup(groupName string, fishCount int64, from int64, till int64) ([]domain.FishSpecies, error)
	FindByName(groupName string) (domain.SensorGroup, error)
}

type groupService struct {
	groupRepo      database.GroupRepository
	sensorService  SensorService
	readingService ReadingService
	fishService    FishSpeciesService
}

func NewGroupService(r database.GroupRepository, ss SensorService, rs ReadingService, fs FishSpeciesService) GroupService {
	return groupService{
		groupRepo:      r,
		sensorService:  ss,
		readingService: rs,
		fishService:    fs,
	}
}

func (s groupService) GetAverageTemperatureForGroup(groupName string) (float64, error) {
	group, err := s.FindByName(groupName)
	if err != nil {
		return 0, err
	}

	sensors, err := s.sensorService.FindByGroupID(group.ID)
	if err != nil {
		return 0, err
	}

	var sensorsIDs []int64
	for _, v := range sensors {
		sensorsIDs = append(sensorsIDs, v.ID)
	}

	return s.readingService.GetAverageTemperatureForGroup(sensorsIDs)
}

func (s groupService) GetAverageTransparencyForGroup(groupName string) (float64, error) {
	group, err := s.FindByName(groupName)
	if err != nil {
		return 0, err
	}

	sensors, err := s.sensorService.FindByGroupID(group.ID)
	if err != nil {
		return 0, err
	}

	var sensorsIDs []int64
	for _, v := range sensors {
		sensorsIDs = append(sensorsIDs, v.ID)
	}

	return s.readingService.GetAverageTransparencyForGroup(sensorsIDs)
}

func (s groupService) GetFishesForGroup(groupName string) ([]domain.FishSpecies, error) {
	group, err := s.FindByName(groupName)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	sensors, err := s.sensorService.FindByGroupID(group.ID)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	var sensorsIDs []int64
	for _, v := range sensors {
		sensorsIDs = append(sensorsIDs, v.ID)
	}

	readings, err := s.readingService.FindBySensorsIDs(sensorsIDs, 0, 0)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	var readingsIDs []int64
	for _, v := range readings {
		readingsIDs = append(readingsIDs, v.ID)
	}

	return s.fishService.GetFishesForGroup(readingsIDs)
}

func (s groupService) GetTopFishesForGroup(groupName string, fishCount int64, from int64, till int64) ([]domain.FishSpecies, error) {
	group, err := s.FindByName(groupName)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	sensors, err := s.sensorService.FindByGroupID(group.ID)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	var sensorsIDs []int64
	for _, v := range sensors {
		sensorsIDs = append(sensorsIDs, v.ID)
	}

	readings, err := s.readingService.FindBySensorsIDs(sensorsIDs, from, till)
	if err != nil {
		return []domain.FishSpecies{}, err
	}

	var readingsIDs []int64
	for _, v := range readings {
		readingsIDs = append(readingsIDs, v.ID)
	}

	return s.fishService.GetTopFishesForGroup(readingsIDs, fishCount)
}

func (s groupService) FindByName(groupName string) (domain.SensorGroup, error) {
	return s.groupRepo.FindByName(groupName)
}
