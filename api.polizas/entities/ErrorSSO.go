package entities

type Error struct {
	ErrorCode   int    `json:"errorCode"`
	UserMessage string `json:"userMessage"`
}
