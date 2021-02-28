package middlewares

import (
	"fmt"
	"log"
	"mekarTechnicalTest/utils/helper"
	"net/http"
)

//APILog is a middleware for record accessed route
func APILog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		logData := fmt.Sprintf("accessing path %v with application %v from %v", r.RequestURI, userAgent, r.RemoteAddr)
		log.Printf(" | %v", logData)
		helper.LogApp(logData)
		next.ServeHTTP(w, r)
	})
}
