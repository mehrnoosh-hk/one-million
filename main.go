package main

import (
	"app/httpADAPTER"
	"log"
	"net/http"
)

func main() {
	mux := httpADAPTER.CreateServer()
	log.Fatal(http.ListenAndServe(":9990", mux))
}
