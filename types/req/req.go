package req

type DemoDetailReq struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required,min=1"`
}

type DemoListReq struct{}

type DemoUpdateReq struct{}

type DemoAddReq struct{}

type DemoDeleteReq struct{}

type DemoCreateReq struct{}
