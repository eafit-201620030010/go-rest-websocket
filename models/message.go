package models

type WebsocketMessage struct {
	Type    string      `json:"created_at"`
	Payload interface{} `json:"payload"`
}
