package main

import (
	"encoding/base64"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":5005", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		// remove basic
		auth = auth[6:]
		log.Println(auth)

		// decode
		raw, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(raw)

		log.Println(string(raw))
	}))

	if err != nil {
		log.Fatal(err)
	}
}
