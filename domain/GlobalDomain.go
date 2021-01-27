package domain

type Result struct {
	Status int16    `json:"status"`
	Msg    string   `json:"msg"`
	Data   struct{} `json:"data"`
}

type ActionInfo struct {
	Action string   `json:"action"`
	Data   interface{} `json:"data"`
}

