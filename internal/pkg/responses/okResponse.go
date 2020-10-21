package responses

import (
	"encoding/json"
	"net/http"
)

func SendDataResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(HttpResponse{Data: data, Code: 200})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
}

func SendOkResponse(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(HttpResponse{Code: 200})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
}
