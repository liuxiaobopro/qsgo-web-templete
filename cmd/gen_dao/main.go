package main

import (
	"fmt"

	xormx "github.com/liuxiaobopro/gobox/xorm"
	"xorm.io/xorm"
)

func main() {
	engine, err := xorm.NewEngine("mysql", "root:111111@tcp(192.168.3.74:3306)/demo?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	gxd := xormx.NewGenXormDao(
		xormx.WithMysql(engine),
		xormx.WithDaoMysqlPath("dao/mysql"),
		xormx.WithProject("qsgo-web-templete"),
		xormx.WithPrefix("qsgo_"),
		xormx.WithProgramTemplatePath("./tpl/dao_default.tpl"),
		xormx.WithDefaultTemplatePath("./tpl/dao_program.tpl"),
	)
	if err := gxd.Gen(); err != nil {
		panic(err)
	}
	if err := engine.Close(); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
