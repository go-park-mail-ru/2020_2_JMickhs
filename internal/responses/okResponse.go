package responses

import (
	"encoding/json"
	"net/http"
)

func SendOkResponse(w http.ResponseWriter, data interface{} ){
	err := json.NewEncoder(w).Encode(HttpResponse{Data: data})
	if err != nil {
		SendErrorResponse(w,http.StatusInternalServerError,err)
	}
}