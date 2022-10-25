package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func main() {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "5000")
	addr := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/ping", pingHandler)
	log.Printf("starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("could not start server at %s due to %s", addr, err)
	}

}
