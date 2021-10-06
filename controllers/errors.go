package controllers

import (
	"net/http"
	errors "project/business"
)

func ErrorCode(err error) (code int) {
	if err == errors.ErrAPIFound {
		code = http.StatusBadRequest
	} else if err == errors.ErrDuplicateData {
		code = http.StatusBadRequest
	} else if err == errors.ErrFillData {
		code = http.StatusBadRequest
	} else if err == errors.ErrInternalServer {
		code = http.StatusInternalServerError
	} else if err == errors.ErrMovieResource {
		code = http.StatusNoContent
	} else if err == errors.ErrNotFound {
		code = http.StatusNoContent
	} else if err == errors.ErrNotProFound {
		code = http.StatusForbidden
	} else if err == errors.ErrPayNotFound {
		code = http.StatusNoContent
	} else if err == errors.ErrSubsNotFound {
		code = http.StatusNoContent
	} else if err == errors.ErrUserResource {
		code = http.StatusNoContent
	} else if err == errors.ErrUsernamePasswordNotFound {
		code = http.StatusBadRequest
	}
	return code
}
