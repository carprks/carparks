package healthcheck

type Health struct {
	Name         string   `json:"name"`
	URL          string   `json:"url"`
	Status       string   `json:"status"`
	Dependencies []Health `json:"dependencies,omitempty"`
}

type Dependencies struct {
	Dependencies []Depedency `json:"dependencies"`
}
type Depedency struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Ping bool   `json:"ping"`
}
