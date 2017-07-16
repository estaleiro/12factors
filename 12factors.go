package main

import (
	"log"
	"net/http"
	"strings"

	factors "github.com/rikatz/12factors/factors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var rc int
	var msg string
	switch factor := strings.Split(r.URL.Path, "/"); factor[1] {
	case "factor3":
		rc, msg = factors.Factor3()
	case "factor6":
		rc, msg = factors.Factor6(factor, w, r)
	default:
		rc, msg = http.StatusOK, "Hello, this is 12 Factor demonstration"
	}
	w.WriteHeader(rc)
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
}
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
