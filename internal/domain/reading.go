package domain

import "time"

type Reading struct {
	ID           int64     `json:"id"`
	SensorID     int64     `json:"sensor_id"`
	Timestamp    time.Time `json:"timestamp"`
	Temperature  float64   `json:"temperature"`
	Transparency int64     `json:"transparency"`
	//FishSpecies  []FishSpecies `json:"fish_species"`
}

//type FishSpecies struct {
//	Name  string `json:"name"`
//	Count int64  `json:"count"`
//}
