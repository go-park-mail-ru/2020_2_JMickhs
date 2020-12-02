package responses

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(HttpResponse{Code: code})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
