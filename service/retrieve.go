package service

import "net/http"

// RetrieveList gets a list of carparks for the owner
func RetrieveList(w http.ResponseWriter, r *http.Request){
	c := CarParkOwner{
		ID: "tester",
	}
	c.ListCarParks()
}

// Retrieve get the carpark details
func Retrieve(w http.ResponseWriter, r *http.Request) {

}

// ListCarParks list the carparks
func (c CarParkOwner) ListCarParks() ([]CarPark, error) {
	r := []CarPark{}

	return r, nil
}
