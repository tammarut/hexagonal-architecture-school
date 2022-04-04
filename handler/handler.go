package handler

import (
	"code-bangkok/errs"
	"fmt"
	"net/http"
)

func handleError(writer http.ResponseWriter, err error) {
	switch myErr := err.(type) {
	case errs.AppError:
		writer.WriteHeader(myErr.Code)
		fmt.Fprintln(writer, myErr)
	default:
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, myErr)
	}
}
