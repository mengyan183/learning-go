package built_in_package

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// 内置包 strings

func TestStrings(t *testing.T) {
	s := "Hello World"
	// 是否包含指定字符串
	assert.True(t, strings.Contains(s, "ld"))
	// 是否包含任意字符串中的任意单一字符 字符串
	assert.True(t, strings.ContainsAny(s, "wd"))
	// 指定字符出现次数
	assert.Equal(t, 3, strings.Count(s, "l"))
	// 开头
	assert.True(t, strings.HasPrefix(s, "H"))
	// 结尾
	assert.True(t, strings.HasSuffix(s, "ld"))
	// 查找指定字符串存在的索引位置
	assert.Equal(t, strings.Index(s, "ll"), 2)
	// 查找任意字符出现在字符串中的位置, 这里返回的是任意一个字符第一次匹配到索引
	assert.Equal(t, strings.IndexAny(s, "ol"), 2)
	// 查找指定字符串在字符串中的最后一个索引位置
	assert.Equal(t, strings.LastIndex(s, "l"), 9)
	// 字符串切割
	assert.Equal(t, strings.Split(s, " "), []string{"Hello", "World"})
	// 字符串拼接
	assert.Equal(t, strings.Join([]string{"Hello", "World"}, " "), s)
	// 字符串替换, 最后一个参数表示要替换的字符个数
	t.Log(strings.Replace(s,"H","h",1))
	// 替换全部匹配到的字符
	t.Log(strings.ReplaceAll(s,"l","L"))

}
