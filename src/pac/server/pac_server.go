//package server
// 包名可以和当前文件路径名不保持一致,但要求相同目录下的文件的包名必须一致
package server_pac

import "fmt"

func privateFunc() {
	fmt.Println("私有方法,不允许包外引用")
}

func PublicFunc() {
	fmt.Println("公共方法")
}
//
//func ImportClient()  {
//	client.Client()
//}
