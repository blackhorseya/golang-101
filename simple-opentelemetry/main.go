package main

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/rolldice", rolldice)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // Use default multiplexer
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func rolldice(w http.ResponseWriter, r *http.Request) {
	n, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		log.Printf("Write failed: %v\n", err)
		return
	}
	roll := n.Int64() + 1

	resp := strconv.Itoa(int(roll)) + "\n"
	_, err = io.WriteString(w, resp)
	if err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
