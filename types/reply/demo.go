package reply

type DemoGetReply struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type DemoPostReply struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
