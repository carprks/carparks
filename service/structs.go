package service

// CreateModel standard to check the create against
type CreateModel struct {
	// Inputs
	Name      string `json:"name"`
	Spaces    int    `json:"spaces"`
	PostCode  string `json:"postcode"`

	// Model
	Longitude string
	Latitude  string
}

// Response did it work or not
type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
