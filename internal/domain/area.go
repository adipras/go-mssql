package domain

// Response represents the api respone
type Response struct {
	Message string `json:"message"`
}

// List Area
type Areas struct {
	AreaID   string  `json:"ar_id"`
	AreaCode string  `json:"ar_code"`
	AreaName string  `json:"ar_name"`
	AreaAW   string  `json:"ar_aw"`
	AreaOC   string  `json:"ar_oc"`
	AreaLat  float32 `json:"ar_latitude"`
	AreaLong float32 `json:"ar_longitude"`
}

// area repository implements
type AreaRepository interface {
	ListAreas() (*[]Areas, error)
}
