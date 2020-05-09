package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	helpers "proj1/helpers"

	"github.com/joho/godotenv"
)

var hmacSampleSecret []byte

//GenerateAuthToken - Generates Auth Session
func GenerateAuthToken(w http.ResponseWriter, r *http.Request) {
	tokenString, err := helpers.GenerateAuthToken(hmacSampleSecret)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, tokenString)
}

//ValidateSession - Validates Auth Session
func ValidateSession(w http.ResponseWriter, r *http.Request) {
	// In Realtime Token is added to Authorization Header
	var session struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, "Session Details Not Found", http.StatusBadRequest)
		return
	}

	response, valError := helpers.ValidateSession(session.Token, hmacSampleSecret)
	if valError != nil {
		http.Error(w, valError.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, response)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hmacSampleSecret = []byte(os.Getenv("SECRET"))
}

func handleHTTP() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", GenerateAuthToken)
	mux.HandleFunc("/validate", ValidateSession)
	log.Fatal(http.ListenAndServe(":8081", mux))
}
func main() {
	fmt.Println("GO Runnning")
	handleHTTP()
}
