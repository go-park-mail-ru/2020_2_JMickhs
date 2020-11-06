//go:generate easyjson -all paginationModel.go
package paginationModel


// easyjson:json
type PaginationInfo struct {
	NextLink string `json:"next,omitempty"`
	PrevLink string `json:"prev,omitempty"`
	ItemsCount int `json:"items_count,omitempty"`
}
