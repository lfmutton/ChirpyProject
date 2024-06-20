package main

import (
	"net/http"
	"log"
)

func main(){
	const port = "8080"
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}