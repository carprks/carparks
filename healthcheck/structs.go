package healthcheck

// Health the result to see if it works
type Health struct {
	Name         string   `json:"name"`
	URL          string   `json:"url"`
	Status       string   `json:"status"`
	Dependencies []Health `json:"dependencies,omitempty"`
}

// Dependencies the list of dependencies to test
type Dependencies struct {
	Dependencies []Dependency `json:"dependencies"`
}

// Dependency to test
type Dependency struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Ping bool   `json:"ping"`
}
