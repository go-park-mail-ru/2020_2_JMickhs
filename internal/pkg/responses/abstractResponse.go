package responses

type HttpResponse struct {
	Data interface{} `json:"data,omitempty"`
	Code int         `json:"code"`
}
