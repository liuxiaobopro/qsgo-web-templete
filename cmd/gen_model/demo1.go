package models

import (
	"time"
)

type Demo1 struct {
	Id       int       `xorm:"not null pk autoincr INT" json:"id"`
	Name     string    `xorm:"VARCHAR(255)" json:"name"`
	CreateAt time.Time `xorm:"comment('创建时间') DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"comment('更新时间') DATETIME" json:"update_at"`
}

func (m *Demo1) TableComment() string {
	return "demo1"
}
