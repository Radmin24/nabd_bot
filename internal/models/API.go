package models

type ControllerResponce struct {
	Answer    string   `json:"answer"`
	Delay     int      `json:"delay"`
	Keyboard  Keyboard `json:"keyboard"`
	IsKb      bool     `json:"isKb"`
	Image     string   `json:"image"`
	IsNextMsg bool     `json:"isNextMsg"`
	Id        int      `json:"id"`
}

type Keyboard struct {
	Button []Button `json:"buttons"`
	Type   string   `json:"type"`
}

type Button struct {
	Caption string `json:"caption"`
	Data    string `json:"data"`
	Order   int    `json:"order"`
	Row     int    `json:"row"`
}

type gRPCMessage struct {
	Msg      string   `json:"mes"`
	Delay    int      `json:"delay"`
	Keyboard Keyboard `json:"keyboard"`
	IsKb     bool     `json:"isKb"`
	Image    string   `json:"image"`
	ChatId   int64    `json:"chat_id"`
}
