package service

import (
	"fmt"
	"net/http"
)

// RetrieveList gets a list of carparks for the owner
func RetrieveList(w http.ResponseWriter, r *http.Request) {
	c := CarParkOwner{
		ID: "tester",
	}
	c.ListCarParks()

	fmt.Println("Retrieve List")
}

// Retrieve get the carpark details
func Retrieve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieve")
}

// ListCarParks list the carparks
func (c CarParkOwner) ListCarParks() ([]CarPark, error) {
	r := []CarPark{}

	return r, nil
}
