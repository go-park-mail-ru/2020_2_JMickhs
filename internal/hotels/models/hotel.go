package models

type Hotel struct {
	HotelID     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Location    string `json:"location"`
	Rating      string `json:"rating"`
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

type SearchData struct {
	Hotels []Hotel `json:"hotels"`
	Cursor Cursor  `json:"cursor"`
}

type Cursor struct {
	NextCursor string `json:"nextcursor"`
	PrevCursor string `json:"prevcursor"`
}

type FilterData struct {
	Rating string
	ID     string
}
