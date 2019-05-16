package main

import (
	"fmt"
	"github.com/carprks/carparks/healthcheck"
	"github.com/carprks/carparks/probe"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func _main(args []string) int {
	if os.Getenv("GOOGLE_API_KEY") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("Env Load: %v", err))
			return 0
		}
	}

	if len(os.Getenv("GOOGLE_API_KEY")) == 0 {
		fmt.Println("Cant load google api key")
		return 0
	}

	port := "80"
	if len(os.Getenv("PORT")) > 2 {
		port = os.Getenv("PORT")
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Probe
	router.Get("/probe", probe.HTTP)

	// HealthCheck
	router.Get("/healthcheck", healthcheck.HTTP)

	// Wildcard
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Tester"))
		if err != nil {
			fmt.Println(fmt.Errorf("write wild err: %v", err))
			return
		}
		fmt.Println(fmt.Sprintf("Request: %v", r))
	})

	// Router
	// router := mux.NewRouter()
	//
	// // Probe
	// router.HandleFunc("/probe", probe.HTTP) // needed to keep healthy
	// router.HandleFunc(fmt.Sprintf("%s/probe", os.Getenv("SITE_PREFIX")), probe.HTTP)
	//
	// // Health Check
	// router.HandleFunc(fmt.Sprintf("%s/healthcheck", os.Getenv("SITE_PREFIX")), healthcheck.HealthCheck)
	//
	// // Service
	// router.HandleFunc(fmt.Sprintf("%s/", os.Getenv("SITE_PREFIX")), service.Create).Methods("POST")
	// router.HandleFunc(fmt.Sprintf("%s/", os.Getenv("SITE_PREFIX")), service.RetrieveList).Methods("GET")
	// router.HandleFunc(fmt.Sprintf("%s/{carparkId:[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}}", os.Getenv("SITE_PREFIX")), service.Retrieve).Methods("GET")
	// router.HandleFunc(fmt.Sprintf("%s/{carparkId:[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}}", os.Getenv("SITE_PREFIX")), service.Patch).Methods("PATCH")

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
