package apiserver

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter // if we use anonim field we dont have to implement all responsewriter methods
	code int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}