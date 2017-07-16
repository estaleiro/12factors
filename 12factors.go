package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	factors "github.com/rikatz/12factors/factors"
)

// ServerHealth determines the Server Health Status
var ServerHealth = &factors.ServerHealth{
	Rc:  http.StatusOK,
	Msg: "O serviço está operando normalmente",
}

func handler(w http.ResponseWriter, r *http.Request) {
	var rc int
	var msg string

	switch factor := strings.Split(r.URL.Path, "/"); factor[1] {
	case "factor3":
		rc, msg = factors.Factor3()
	case "factor6":
		rc, msg = factors.Factor6(factor, w, r)
	case "factor9":
		rc, msg = factors.Factor9(&ServerHealth, factor)
	default:
		rc, msg = http.StatusOK, "Hello, this is 12 Factor demonstration"
	}
	w.WriteHeader(rc)
	w.Write([]byte(msg))
	hostname, _ := os.Hostname()
	log.SetFlags(0)
	t := time.Now()
	log.Printf("[%s] %s %s %s %s %s %d\n", t.Format("02/01/2006 15:04:05"), strings.Split(r.RemoteAddr, ":")[0], hostname, r.Method, r.URL.Path, r.Proto, rc)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
