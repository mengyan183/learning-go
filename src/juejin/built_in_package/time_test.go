package built_in_package

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(now)
	// 要求模板中的日期时间 必须为 2006-01-02 15:04:05 , 否则就无法转换正确的时间
	t.Log(now.Format("2006-01-02"))
	// 字符串转换为时间
	t.Log(time.Parse("2006-01-02", "2020-08-02"))
	// 对当前时间增加一个小时
	nt := now.Add(time.Hour)
	// 获取时间戳
	i := nt.Unix()
	// 获取时间差值
	d := now.Sub(nt) // now - nt
	t.Log(d)
	// 将时间戳转换为time并格式化字符串
	s := time.Unix(i,0).Format("2006-01-02")
	t.Log(s)
}
