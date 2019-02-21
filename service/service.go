package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Create handle post
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createModel := CreateModel{}
	err := json.NewDecoder(r.Body).Decode(&createModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Invalid Model: %v", err)
		return
	}

	response, err := createModel.CreateCarPark()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(Response{
			Response: response,
			Error:    err.Error(),
		})
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Response: "",
				Error:    err.Error(),
			})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Response: response,
		Error:    "",
	})
}

func (c CreateModel) CreateCarPark() (string, error) {
	if c.Spaces == 0 {
		return "failure", errors.New("no spaces")
	}

	if len(c.Name) == 0 {
		return "failure", errors.New("name missing")
	}

	return "success", nil
}
