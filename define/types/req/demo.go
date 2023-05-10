package req

type DemoGetReq struct {
	Id int `json:"id" form:"id"`
}

type DemoPostReq struct {
	Id int `json:"id" form:"id"`
}
