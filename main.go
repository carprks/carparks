package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"main/probe"
	"main/service"
	"net/http"
	"os"
)

func _main(args []string) int {
	if len(os.Getenv("GOOGLE_API_KEY")) == 0 {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Env Load: %v", err)
			return 0
		}

		if len(os.Getenv("GOOGLE_API_KEY")) == 0 {
			log.Printf("Cant load google api key")
			return 0
		}
	}

	port := "80"
	if len(os.Getenv("PORT")) > 2 {
		port = os.Getenv("PORT")
	}

	// router
	router := mux.NewRouter().StrictSlash(true)

	// Create
	router.HandleFunc("/create", service.Create).Methods("POST")

	// Probe
	router.HandleFunc("/probe", probe.Probe)
	router.HandleFunc("/", probe.Probe)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Println("HTTP", err)
		return 1
	}

	log.Println("Died but nicely")
	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}
