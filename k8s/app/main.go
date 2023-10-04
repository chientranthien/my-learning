package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		log.Printf("received from: %v", request.RemoteAddr)
		fmt.Fprintf(w, "hello: %v", request.RemoteAddr)
	})
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatalf("err=%v",err)
		return
	}
}
