package responses

import (
	"encoding/json"
	"net/http"

	"github.com/mailru/easyjson"
)

func SendDataResponse(w http.ResponseWriter, data interface{}) {
	response := HttpResponse{Data: data, Code: 200}
	_, _, err := easyjson.MarshalToHTTPResponseWriter(response, w)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}

func SendOkResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(HttpResponse{Code: 200})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
}
