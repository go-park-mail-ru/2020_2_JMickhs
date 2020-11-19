package customerror

import (
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/package/responses"
)

func PostError(w http.ResponseWriter, req *http.Request, log *logger.CustomLogger, err error, code interface{}) {
	if code != nil {
		err = NewCustomError(err, code.(int), 2)
	}

	log.LogError(req.Context(), err)
	responses.SendErrorResponse(w, StatusCode(ParseCode(err)))
}
