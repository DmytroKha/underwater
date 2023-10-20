package resources

import "github.com/DmytroKha/underwater/internal/domain"

type FishDto struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func (d FishDto) DomainToDto(fs domain.FishSpecies) FishDto {
	return FishDto{
		Name:  fs.Name,
		Count: fs.Count,
	}
}

func (d FishDto) DomainToDtoCollection(f []domain.FishSpecies) []FishDto {

	result := make([]FishDto, len(f))

	for i := range f {
		result[i] = d.DomainToDto(f[i])
	}

	return result
}
