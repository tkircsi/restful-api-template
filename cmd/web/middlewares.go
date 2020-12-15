package main

import "net/http"

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// not used currently. I used chi's logging middleware
// func (app *application) logRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
// 		next.ServeHTTP(w, r)
// 	})
// }
