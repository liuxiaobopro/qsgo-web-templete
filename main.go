package main

import (
	"qsgo-web-templete/initialize"
)

func main() {
	initialize.Log()
	initialize.Config()
	initialize.Http()
}
