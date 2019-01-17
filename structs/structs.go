package structs

type ClientResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ClientRequest struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
type GameState struct {
	Cells []Cell `json:"cells"`
	Food  []Food `json:"food"`
}

type Cell struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Alive bool `json:"alive"`

	Size  int    `json:"size"`
	Kills int    `json:"kills"`
	Meals []Cell `json:"meals"`

	Pos []int `json:"pos"`
}

type Food struct {
	Id    int   `json:"id"`
	Pos   []int `json:"pos"`
	Value int   `json:"value"`
	Alive bool  `json:"alive"`
}
