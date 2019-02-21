package service

// CreateModel standard to check the create against
type CreateModel struct {
	Name      string `json:"name"`
	Spaces    int    `json:"spaces"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
