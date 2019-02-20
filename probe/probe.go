package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handle func(w http.ResponseWriter, req *http.Request) error

func Probe(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		log.Println("Probe Request", string(buf))
	}

	// send status
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(Healthy{
		Status: "pass",
	})
	w.Write(j)

	return
}

func (h Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover", r)
		}
	}()

	if err := h(w, r); err != nil {
		log.Println("Error", err)

		if httpErr, ok := err.(Error); ok {
			http.Error(w, httpErr.Message, httpErr.Code)
		}
	}
}

func (e Error) Error() string {
	if e.Message == "" {
		e.Message = http.StatusText(e.Code)
	}
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}