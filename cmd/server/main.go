package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	} // слушаем все интерфейсы на 8080

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Docker + Compose!")
	})

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
