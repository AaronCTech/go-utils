package utils

import (
	"crypto/rand"
	"fmt"
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

//PrepErrorResponse returns a server error
func PrepErrorResponse(w http.ResponseWriter, err error, code int, data string) error {
	Handlerr(err, data)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(data))

	return err
}

//GenerateUUID generates a very basic UUID
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
