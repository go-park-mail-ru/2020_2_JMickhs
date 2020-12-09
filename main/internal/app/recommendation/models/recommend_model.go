//go:generate easyjson -all recommend_model

package recommModels

import "time"

// easyjson:json
type HotelRecommend struct {
	HotelID  int    `json:"hotel_id" db:"hotel_id"`
	Name     string `json:"name" db:"name"`
	Image    string `json:"image" db:"concat"`
	Location string `json:"location" db:"location"`
	Rating   string `json:"rating" db:"curr_rating"`
}

// easyjson:json
type Hotels struct {
	Hotels []HotelRecommend `json:"hotels"`
}

type RecommendMatrixRow struct {
	UserID   float64 `db:"user_id"`
	RatingID float64 `db:"rating1"`
	HotelID  float64 `db:"hotel1"`
}

type Match struct {
	UserID      int
	Coefficient float64
}

type BestProduct struct {
	HotelID     int
	Coefficient float64
}

type Recommendation struct {
	UserID   int     `db:"user_id"`
	HotelIDs []int64 `db:"h"`
	Time     time.Time
}
