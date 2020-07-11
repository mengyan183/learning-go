package server_pac

//如果包名不一致会提示: Multiple packages in directory: server_pac, server
//package server

func samePacPrivate() {
	// 同路径下的私有方法可以引用
	privateFunc()
}
// 对于相同目录下的文件中不允许有同名的方法
//func PublicFunc() {
//	fmt.Println("公共方法")
//}
// 没有重载
//func PublicFunc(i int) {
//	fmt.Println("公共方法")
//}
// 私有方法也不允许重名
//func privateFunc() {
//	fmt.Println("私有方法,不允许包外引用")
//}