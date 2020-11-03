package customerror

import (
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
)

var convertStatusToHTTP = map[int]int{
	clientError.BadRequest:           http.StatusBadRequest,
	clientError.PaymentReq:           http.StatusPaymentRequired,
	clientError.Locked:               http.StatusLocked,
	clientError.Unauthorizied:        http.StatusUnauthorized,
	clientError.Conflict:             http.StatusConflict,
	clientError.Forbidden:            http.StatusForbidden,
	clientError.Gone:                 http.StatusGone,
	clientError.UnsupportedMediaType: http.StatusUnsupportedMediaType,
	serverError.ServerInternalError:  http.StatusInternalServerError,
	clientError.NotFound: 			  http.StatusNotFound,
}

func StatusCode(code int) int {
	return convertStatusToHTTP[code]
}
