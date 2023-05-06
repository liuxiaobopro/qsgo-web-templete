package models

import (
	typesx "github.com/liuxiaobopro/gobox/types"
)

type Demo struct {
	Id                    int    `xorm:"not null pk autoincr INT" json:"id"`
	Name                  string `xorm:"VARCHAR(255)" json:"name"`
	typesx.ModelCUExtends `xorm:"extends"`
}

func (m *Demo) TableComment() string {
	return "demo"
}
