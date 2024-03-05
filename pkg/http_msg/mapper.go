package httpmsg

import (
	"errors"
	"github.com/tonet-me/tonet-core/logger"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"log/slog"
	"net/http"
)

const (
	internalStatus = 500
)

func Error(err error) (message string, code int) {
	const op = richerror.OP("httpmsg.Error")
	re := new(richerror.RichError)
	if !errors.As(err, &re) {
		return err.Error(), http.StatusBadRequest
	}
	msg := re.Message()
	code = mapKindToHTTPStatusCode(re.Kind())
	// we should not expose unexpected error messages
	if code >= internalStatus {
		logger.GetLogger().Error(string(op), slog.String("internal error", msg))
		msg = errmsg.ErrorMsgSomethingWentWrong
	}

	return msg, code
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.ErrKindInvalid:

		return http.StatusUnprocessableEntity
	case richerror.ErrKindNotFound:

		return http.StatusNotFound
	case richerror.ErrKindForbidden:

		return http.StatusForbidden
	case richerror.ErrKindUnExpected:

		return http.StatusInternalServerError
	case richerror.ErrKindStatusConflict:
		return http.StatusConflict

	default:

		return http.StatusBadRequest
	}
}
