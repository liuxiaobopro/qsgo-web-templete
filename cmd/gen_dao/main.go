package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	xormx "github.com/liuxiaobopro/gobox/xorm"
	"xorm.io/xorm"
)

func main() {
	engine, err := xorm.NewEngine("mysql", "qsgo-web-templete:pxe6iXMNxCj4hf4B@tcp(127.0.0.1:3306)/qsgo-web-templete?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	gxd := xormx.NewGenXormDao(
		xormx.WithMysql(engine),
		xormx.WithDaoMysqlPath("../../dao/mysql"),
		xormx.WithProject("qsgo-web-templete"),
		xormx.WithPrefix(""),
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
