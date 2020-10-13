package responses

type HttpError struct {
	Code int `json:"code"`
}

type HttpResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *HttpError  `json:"error,omitempty"`
}
