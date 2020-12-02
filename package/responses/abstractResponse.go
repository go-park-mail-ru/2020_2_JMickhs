//go:generate  easyjson -all abstractResponse.go
package responses

// easyjson:json
type HttpResponse struct {
	Data interface{} `json:"data,omitempty"`
	Code int         `json:"code"`
}
