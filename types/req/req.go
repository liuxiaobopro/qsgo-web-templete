package req

type DemoDetailReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type DemoListReq struct{}

type DemoUpdateReq struct{}

type DemoAddReq struct{}

type DemoDeleteReq struct{}


type DemoCreateReq struct {}
	