package responses

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, code int, err error){
	httpErr := HttpError{
		Code: code,
		Message: err.Error(),
	}

	err = json.NewEncoder(w).Encode(HttpResponse{Error: &httpErr})
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
}


