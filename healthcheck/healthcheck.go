package healthcheck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Name: os.Getenv("SERVICE_NAME"),
		URL: r.Host,
		Status: "pass",
	}

	// Dependencies
	depsString := os.Getenv("SERVICE_DEPENDENCIES")
	deps := Dependencies{}
	err := json.Unmarshal([]byte(depsString), &deps)
	if err != nil {
		w.Header().Set("Content-Type", "application/health+json")
		health.Status = "fail"
		j, _ := json.Marshal(health)
		w.Write(j)
		w.WriteHeader(http.StatusOK)

		return
	}

	for _, dep := range deps.Dependencies {
		h := Health{}

		if dep.Ping {
			h = Health{
				Name: dep.Name,
				URL: dep.URL,
				Status: "fail",
			}

			d, _ := http.Get(dep.URL)
			if d.StatusCode != 500 {
				h.Status = "pass"
			}
		} else {
			durl := fmt.Sprintf("https://%s/healthcheck", dep)
			d, err := http.Get(durl)
			if err != nil {
				h = Health{
					URL:    durl,
					Status: "fail",
				}
			}

			b, err := ioutil.ReadAll(d.Body)
			if err != nil {
				h = Health{
					URL:    durl,
					Status: "fail",
				}
			}
			json.Unmarshal(b, &h)
		}

		health.Dependencies = append(health.Dependencies, h)
	}

	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(health)
	w.Write(j)
	w.WriteHeader(http.StatusOK)

	return
}