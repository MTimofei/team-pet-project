package myerr

import (
	"log"
	"net/http"
)

func ServesError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
