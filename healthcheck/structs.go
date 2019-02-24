package healthcheck

type Health struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Status string `json:"status"`
	Dependencies []Health `json:"dependencies,omitempty"`
}