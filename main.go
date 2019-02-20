package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"main/probe"
)

func _main(args []string) int {
	port := "80"
	if len(os.Getenv("PORT")) > 2 {
		port = os.Getenv("PORT")
	}

	// router
	router := mux.NewRouter().StrictSlash(true)

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
