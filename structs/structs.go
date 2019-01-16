package structs

type WsMsg struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
