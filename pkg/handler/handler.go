package handler

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/seregayoga/factorial/pkg/factorial"
)

// CalculateHandler handles calculate requests
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	n, err := getNumber(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := factorial.Calculate(n)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, res.String())
}

func getNumber(r *http.Request) (int64, error) {
	ns := r.URL.Query().Get("n")
	if ns == "" {
		return 0, errors.New("Missing \"n\" parameter in query")
	}

	n, err := strconv.ParseInt(ns, 10, 64)
	if err != nil || n < 0 {
		return 0, errors.New("Parameter \"n\" should be a positive number")
	}

	return n, nil
}
