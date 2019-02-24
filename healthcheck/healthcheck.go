package healthcheck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dep := os.Getenv("SERVICE_DEPENDENCIES")
	deps := strings.Split(dep, ",")

	health := Health{
		Name: os.Getenv("SERVICE_NAME"),
		URL: r.Host,
		Status: "pass",
	}

	if len(deps) >= 1 && deps[0] != "" {
		for _, dep := range deps {
			durl := fmt.Sprintf("https://%s/healthcheck", dep)

			h := Health{}

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

			health.Dependencies = append(health.Dependencies, h)
		}
	}

	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(health)
	w.Write(j)
	w.WriteHeader(http.StatusOK)

	return
}