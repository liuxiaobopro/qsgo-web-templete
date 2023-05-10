package models

import (
	"github.com/liuxiaobopro/gobox/time"
)

type Demo struct {
	Id       int       `xorm:"not null pk autoincr INT" json:"id"`
	Name     string    `xorm:"VARCHAR(255)" json:"name"`
	CreateAt time.Time `xorm:"comment('创建时间') DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"comment('更新时间') DATETIME" json:"update_at"`
}

func (m *Demo) TableComment() string {
	return "demo"
}
