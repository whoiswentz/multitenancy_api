package controllers

import "net/http"

func LogiController(w http.ResponseWriter, r *http.Request) {
	h := r.Host
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h))
}
