package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 执行shell命令
	cmd := exec.Command("reverse", "-f", "reverse.yml")
	// 获取输出对象，可以从该对象中读取输出结果
	opBytes, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	// 打印输出结果
	print(string(opBytes))
	fmt.Println("Done!")
}
