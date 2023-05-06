package models

import (
	"time"
)

type Demo struct {
	Id       int       `xorm:"not null pk autoincr INT" json:"id"`
	CreateAt time.Time `xorm:"DATETIME" json:"create_at"`
	Name     string    `xorm:"VARCHAR(255)" json:"name"`
}

func (m *Demo) TableComment() string {
	return "demo"
}
