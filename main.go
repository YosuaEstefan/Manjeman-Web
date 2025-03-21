package main

import (
	"log"
	"net/http"
	"tugas-deployment/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
