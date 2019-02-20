package probe

type Error struct {
	Code    int
	Message string
}

type Healthy struct {
	Status string `json:"status"`
}
