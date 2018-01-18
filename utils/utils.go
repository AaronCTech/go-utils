package utils

import (
	"log"
	"net/http"
	"time"
)

//Handlerr logs errors and continues application execution
func Handlerr(err error, text string) error {
	if err != nil {
		log.Printf("%s: %v", text, err)
	}

	return err
}

//Fatalerr logs errors and exits the application
func Fatalerr(err error, text string) error {
	if err != nil {
		log.Fatalf("%s: %v", text, err)
	}

	return err
}

//TimeTrack is a debugging function that allows us to track execution times
func TimeTrack(start time.Time, name string) {
	//start := time.now()
	elapsed := time.Since(start)

	log.Printf("%s took %s", name, elapsed)
}

//PrepResponseWriter prepares a successful response to be returned
func PrepResponseWriter(w http.ResponseWriter, data []byte) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

	return w
}

