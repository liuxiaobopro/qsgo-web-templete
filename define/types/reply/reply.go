package reply

import (
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type DemoDetailReply struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DemoListReply struct {
	*replyx.List
}

type DemoUpdateReply struct{}

type DemoAddReply struct{}

type DemoDeleteReply struct{}

type DemoCreateReply struct{}
