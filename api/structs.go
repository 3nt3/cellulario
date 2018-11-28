package api

type cell struct {
	id   int `json:"id"`
	name int `json:"name"`

	size  float64 `json:"size"`
	kills int     `json:"kills"`
	meals []cell  `json:"meals"`

	pos []int `json:"pos"`
}

type food struct {
	id    int `json:"id"`
	pos   []int `json:"pos"`
	value int `json:"value"`
}
