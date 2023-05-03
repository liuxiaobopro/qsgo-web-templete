package initialize

import (
	"os"
	"qsgo-web-templete/global"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func Db() {
	engine, err := xorm.NewEngine("mysql", global.Conf.Mysql.Dns)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("runtime/log/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(log.NewSimpleLogger(f)) // 设置日志输出位置
	engine.ShowSQL(true)                     // 在控制台打印sql
	engine.SetLogLevel(log.LOG_DEBUG)        // 设置日志级别
	engine.SetMaxIdleConns(10)               // 设置连接池的空闲数大小
	engine.SetMaxOpenConns(100)              // 设置最大打开连接数
	engine.SetConnMaxLifetime(10)            // 设置连接的最大生命周期
	engine.SetTZLocation(time.Local)         // 设置时区

	global.DB = engine
}
