package entities

type MetaSSO struct {
	Status string `json:"status"`
	Count  int
	Error  Error
}
