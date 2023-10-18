package domain

type Sensor struct {
	ID             int64   `json:"id"`
	Codename       string  `json:"codename"`
	GroupID        int64   `json:"group_id"`
	X              float64 `json:"x"`
	Y              float64 `json:"y"`
	Z              float64 `json:"z"`
	DataOutputRate int64   `json:"data_output_rate"`
}
