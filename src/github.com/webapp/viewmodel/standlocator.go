package viewmodel

// StandLocator struct
type StandLocator struct {
	Title  string
	Active string
}

// NewStandLocator constructor
func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
	}
	return result
}

// StandCoordinate struct to represent the stand locations
type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
