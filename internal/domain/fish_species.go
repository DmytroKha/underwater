package domain

type FishSpecies struct {
	ID        int64  `json:"id"`
	ReadingID int64  `json:"reading_id"`
	Name      string `json:"name"`
	Count     int64  `json:"count"`
}

type Fishes struct {
	Items []FishSpecies
	Total uint64
	Pages uint64
}
