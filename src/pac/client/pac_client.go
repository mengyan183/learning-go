package client

import (
	"fmt"
)

func init() {
	fmt.Println("client init1")
}

func Client() {
	fmt.Println("Client")
}

func ImportServer(){
	// 当发生循环引用时,编译时不会报错,执行时会 抛出 can't load package: import cycle not allowed
	//server_pac.ImportClient()
	//server_pac.PublicFunc()
}
