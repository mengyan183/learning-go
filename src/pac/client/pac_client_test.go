package client

import (
	cm "github.com/easierway/concurrent_map" // 引入远程package 可以通过 go get -u 拉取最新的远程代码
	"testing"
)
// 这里引入的是src下的文件路径;并定义别名
import server "pac/server"

func TestImport(t *testing.T) {
	// 真正调用是使用的文件夹下 代码中定义的package名称
	//server_pac.PublicFunc()
	// 当在import 中对引入的package定义别名,可以直接使用别名来访问引入包下的方法,就不用使用文件中定义的package name
	server.PublicFunc()
	// 不能调用不同package下的私有方法
	//server_pac.privateFunc()
	ImportServer()

	c := cm.CreateConcurrentMap(10)
	c.Set(cm.StrKey("key"), 10)
	c.Set(cm.I64Key(10), "10")
	if v, ok := c.Get(cm.StrKey("key")); ok {
		t.Log(v)
	}
	if v, ok := c.Get(cm.I64Key(10)); ok {
		t.Log(v)
	}
}
