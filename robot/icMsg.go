package robot

type IcMsg struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
