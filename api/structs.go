package api

type room struct {
	id    int
	cells []cell
}

type cell struct {
	id   int `json:"id"`
	name int `json:"name"`

	size  float64 `json:"size"`
	kills int     `json:"kills"`
	meals []cell  `json:"meals"`
}
