package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintf(w, e.Message)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, e.Error())
	}
}
