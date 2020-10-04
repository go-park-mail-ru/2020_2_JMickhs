package models

// swagger:response Hotel
type Hotel struct {
	ID          int    `json:id`
	Name        string `json:name`
	Description string `json:description`
	Image       string `json:image`
}
