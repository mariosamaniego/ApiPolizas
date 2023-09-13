package entities

type ResponseSSO struct {
	Meta MetaSSO `json:"meta"`
	Data string  `json:"data"`
}
