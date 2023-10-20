package app

import (
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
)

type GroupService interface {
	GetAverageTemperatureForGroup(groupName string) (float64, error)
	GetAverageTransparencyForGroup(groupName string) (float64, error)
	FindByName(groupName string) (domain.SensorGroup, error)
}

type groupService struct {
	groupRepo      database.GroupRepository
	sensorService  SensorService
	readingService ReadingService
}

func NewGroupService(r database.GroupRepository, ss SensorService, rs ReadingService) GroupService {
	return groupService{
		groupRepo:      r,
		sensorService:  ss,
		readingService: rs,
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

func (s groupService) FindByName(groupName string) (domain.SensorGroup, error) {
	return s.groupRepo.FindByName(groupName)
}
