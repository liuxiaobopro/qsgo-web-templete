package main

import (
	"qsgo-web-templete/initialize"
)

func main() {
	initialize.Config()
	initialize.Log()
	initialize.Db()
	initialize.Dao()
	initialize.Http()
}
