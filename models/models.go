package models

import (
	"github.com/liuxiaobopro/gobox/time"
)

type Demo struct {
	Id       int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Name     string    `xorm:"VARCHAR(255)" json:"name"`
	CreateAt time.Time `xorm:"comment('创建时间') DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"comment('更新时间') DATETIME" json:"update_at"`
}

func (m *Demo) TableComment() string {
	return "demo"
}

type Narration struct {
	Id       int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Content  string    `xorm:"not null default '' comment('内容') unique VARCHAR(500)" json:"content"`
	IsUse    int       `xorm:"not null default 1 comment('是否使用(1使用 2不使用)') TINYINT(1)" json:"is_use"`
	CreateAt time.Time `xorm:"DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"DATETIME" json:"update_at"`
	DeleteAt time.Time `xorm:"DATETIME" json:"delete_at"`
}

func (m *Narration) TableComment() string {
	return "narration"
}

type ReplyLog struct {
	Id       int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Nid      int       `xorm:"not null default 0 comment('话述id') INT(11)" json:"nid"`
	Mode     int       `xorm:"not null default 1 comment('模式') TINYINT(1)" json:"mode"`
	Score    int       `xorm:"not null default 0 comment('分数') INT(11)" json:"score"`
	Name     string    `xorm:"not null default '' comment('名称') VARCHAR(20)" json:"name"`
	Phone    string    `xorm:"not null default '' comment('手机号') VARCHAR(20)" json:"phone"`
	Sex      string    `xorm:"not null default '' comment('性别(-1未知 1男 2女)') VARCHAR(20)" json:"sex"`
	CreateAt time.Time `xorm:"DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"DATETIME" json:"update_at"`
	DeleteAt time.Time `xorm:"DATETIME" json:"delete_at"`
}

func (m *ReplyLog) TableComment() string {
	return "reply_log"
}
