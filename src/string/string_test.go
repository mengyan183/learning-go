package string

import "testing"

func TestString(t *testing.T) {
	var s string // 默认值为空字符串,而不是nil
	t.Log(s, len(s))
	s = "string"
	t.Log(len(s))
	// 虽然string的底层结构是slice,但并并不能对slice中的元素进行更新/新增/删除等操作
	//s[1] = "3" // cannot assign to s[1]
	t.Log(s[0]) // 可以读取字符串中指定索引位置的元素
	// 存储任意二进制数据
	s = "\u6d4b\u8bd5"     // 中文 "测试" 的 Unicode编码
	t.Log(s, len(s))       // 这里的len(s)实际是 当前Unicode编码的个数,而不是对应的中文汉字的长度
	s = "\u6d4b\u8bd5\xFF" // 对于 \xFF 是任意一个二进制数据,而string也是可以直接存储的
	t.Log(s, len(s))

	s = "\xE4\xB8\xA5" //utf8
	t.Log(s, len(s))

}

// 字符串转换为Unicode字符
func TestRune(t *testing.T) {
	s := "中"
	t.Log(len(s))
	// byte 和 rune都是用来表示字符类型的变量类型
	// 不同点在于:
	// byte存储的是uint8的数据,常用来处理ascii字符
	// rune存储的是int32的数据,常用来处理Unicode和utf8 字符
	// rune数据类型
	u := []rune(s)
	t.Log(u, len(u))

	t.Logf("中 Unicode %x", u[0])
	// 字符串是使用UTF8编码进行数据存储; 将20013转换为16进制数据就变为了e4b8ad; 因此实际在使用 切片存储数据时,实际存储的为e4/b8/ad
	t.Logf("中 UTF8 %x", s) //e4b8ad
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// 这里range遍历的实际是 rune切片,而不是byte切片;
	for _, c := range s {
		// %[1]代表获取第一个参数; %c 代表字符;%x代表16进制
		t.Logf("%[1]c %[1]x", c)
	}
}
