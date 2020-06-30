package main

import (
	"fmt"
	"os"
)

// 错误的写法!!!!!!!!!! 会提示 "func main must have no arguments and no return values"; 主方法入口不允许存在返回值
//func main() int {
//	return 1;
//}

func main() {
	fmt.Println("helloWorld")
	// os.Args获取命令行中传递的参数
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		// 自定义输入的参数是从数组下标1开始
		fmt.Println(os.Args[1])
	}
	os.Exit(0) // 程序主动退出,状态值为0代表需要等程序完成结束后才执行程序退出操作;否则则直接退出
}
