package easy_middlware

import (
	"net/http"
	"mime"
	"strings"
)

// Verifies the request Content-Type header and returns a
// StatusUnsupportedMediaType (415) HTTP error response if it's incorrect. The expected
// Content-Type is 'application/json' if the content is non-null. Note: If a charset parameter
// exists, it MUST be UTF-8.
func JsonHeaderCheck(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		mediaType, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		charset, ok := params["charset"]

		if !ok {
			charset = "UTF-8"
		}

		if r.ContentLength > 0 && !(mediaType == "application/json" && strings.ToUpper(charset) == "UTF-8") {

			http.Error(w, "Bad Content-Type or charset, expected 'application/json'", http.StatusUnsupportedMediaType)
			return
		}

		h.ServeHTTP(w, r)
	})
}