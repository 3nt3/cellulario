package structs

type ClientMsg struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
