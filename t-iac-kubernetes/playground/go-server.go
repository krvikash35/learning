package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]string{"message": "ping from go server"})
	w.Write(data)
}
func pingWaitHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2000 * time.Millisecond)
	w.WriteHeader(http.StatusInternalServerError)
	data, _ := json.Marshal(map[string]string{"message": "ping from go server"})
	w.Write(data)
}

func main() {
	host := getEnv("HOST", "0.0.0.0")
	port := getEnv("PORT", "5001")
	addr := fmt.Sprintf("%s:%s", host, port)

	log.Printf("logical cpus: %+v", runtime.NumCPU())
	log.Printf("max procs: %+v", runtime.GOMAXPROCS(0))

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/ping_wait", pingWaitHandler)
	log.Printf("starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("could not start server at %s due to %s", addr, err)
	}

}
