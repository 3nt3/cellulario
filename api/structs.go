package api

type cell struct {
	Id   int
	Name string

	Size  float64
	Kills int
	Meals []cell

	Pos []int
}

type food struct {
	Id    int   `json:"id"`
	Pos   []int `json:"pos"`
	Value int   `json:"value"`
}
