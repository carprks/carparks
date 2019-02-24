package service

// CreateModel standard to check the create against
type CreateModel struct {
	// Inputs
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitemprt"`
	Spaces    int    `json:"spaces,omitempty"`
	PostCode  string `json:"postcode,omitempty"`
	Location  Location `json:"location,omitempty"`
}

// Location the location object
type Location struct {
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
}

// Response did it work or not
type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
