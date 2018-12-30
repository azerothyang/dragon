package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		startTime := time.Now()
		log.Println(r.URL.String())
		next.ServeHTTP(w, r)
		log.Printf("time cost:%d ms\n", time.Since(startTime)/time.Millisecond)
	})
}