package apiserver

import "net/http"

// custon responseWriter type, allows us to have code
type ResponseWriter struct {
	http.ResponseWriter // if we use anonim field we dont have to implement all responsewriter methods
	code int
}

// allows us to take status code, then do origin writeheader method
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}