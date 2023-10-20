package app

import (
	"github.com/DmytroKha/underwater/config"
	"github.com/DmytroKha/underwater/internal/domain"
	"github.com/DmytroKha/underwater/internal/infra/database"
	"math/rand"
	"time"
)

type FishSpeciesService interface {
	Save(readingId int64) error
	GetFishesForGroup(readingsIDs []int64) ([]domain.FishSpecies, error)
	GetTopFishesForGroup(readingsIDs []int64, fishCount int64) ([]domain.FishSpecies, error)
}

type fishSpeciesService struct {
	fishSpeciesRepo database.FishSpeciesRepository
	readingService  ReadingService
}

func NewFishSpeciesService(r database.FishSpeciesRepository) FishSpeciesService {
	return fishSpeciesService{
		fishSpeciesRepo: r,
	}
}

func (s fishSpeciesService) Save(readingId int64) error {
	fishCount := generateFakeFishSpecies()
	for f, c := range fishCount {
		fish := domain.FishSpecies{
			ReadingID: readingId,
			Name:      f,
			Count:     c,
		}
		err := s.fishSpeciesRepo.Save(fish)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s fishSpeciesService) GetFishesForGroup(readingsIDs []int64) ([]domain.FishSpecies, error) {
	return s.fishSpeciesRepo.GetFishesForGroup(readingsIDs)
}

func (s fishSpeciesService) GetTopFishesForGroup(readingsIDs []int64, fishCount int64) ([]domain.FishSpecies, error) {
	return s.fishSpeciesRepo.GetTopFishesForGroup(readingsIDs, fishCount)
}

func generateFakeFishSpecies() map[string]int64 {

	//var fishSpecies []domain.FishSpecies
	numFishSpecies := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	fishCount := make(map[string]int64)

	for i := 0; i < numFishSpecies; i++ {
		randomIndex := rand.Intn(len(config.AllFish))
		randomFish := config.AllFish[randomIndex]
		fishCount[randomFish]++
	}

	return fishCount
}
