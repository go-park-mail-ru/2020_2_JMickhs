package responses

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, code int) {
	httpErr := HttpError{
		Code: code,
	}

	err := json.NewEncoder(w).Encode(HttpResponse{Error: &httpErr})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
