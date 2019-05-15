package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"main/healthcheck"
	"main/probe"
	"main/service"
	"net/http"
	"os"
)

func _main(args []string) int {
	if os.Getenv("GOOGLE_API_KEY") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("Env Load: %v", err))
			return 0
		}

		if len(os.Getenv("GOOGLE_API_KEY")) == 0 {
			fmt.Println("Cant load google api key")
			return 0
		}
	}

	port := "80"
	if len(os.Getenv("PORT")) > 2 {
		port = os.Getenv("PORT")
	}

	// Router
	router := mux.NewRouter()

	// Probe
	router.HandleFunc("/probe", probe.HTTP)
	router.HandleFunc("/carparks/probe", probe.HTTP)

	// Health Check
	router.HandleFunc("/healthcheck", healthcheck.HealthCheck)
	router.HandleFunc("/carparks/healthcheck", healthcheck.HealthCheck)

	// Service
	router.HandleFunc("/", service.Create).Methods("POST")
	router.HandleFunc("/", service.RetrieveList).Methods("GET")
	router.HandleFunc("/{carparkId:[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}}", service.Retrieve).Methods("GET")
	router.HandleFunc("/{carparkId:[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}}", service.Patch).Methods("PATCH")

	// General
	router.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(fmt.Printf("* Err: %v", err))
		}

		fmt.Println(fmt.Printf("Request: %v, %s", r.Body, string(buf)))
	})

	fmt.Println(fmt.Sprintf("Listening on: %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println(fmt.Sprintf("HTTP: %v", err))
		return 1
	}

	fmt.Println("Died but nicely")
	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}
