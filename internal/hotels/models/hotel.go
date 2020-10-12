package models

type Hotel struct {
	HotelID     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Location    string `json:"location"`
}

type HotelPreview struct {
	HotelID  int    `json:"hotel_id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Location string `json:"location"`
}

type SearchString struct {
	Pattern string `json:"pattern"`
}
