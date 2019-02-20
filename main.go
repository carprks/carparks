package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"main/probe"
)

func _main(args []string) int {
	// router
	router := mux.NewRouter().StrictSlash(false)

	// Probe
	router.HandleFunc("/probe", probe.Probe)
	router.HandleFunc("/", probe.Probe)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router); err != nil {
		fmt.Println("HTTP", err)
		return 1
	}

	fmt.Println("Died but nicely")
	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}
