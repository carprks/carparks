package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Probe responds with a healthy model
func Probe(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		log.Println("Probe Request", string(buf))
	}

	// send status
	j, _ := json.Marshal(Healthy{
		Status: "pass",
	})
	w.Header().Set("Content-Type", "application/health+json")
	w.Header().Set("Service", "CarParks")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	return
}
