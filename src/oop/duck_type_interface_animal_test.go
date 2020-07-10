package oop

import "testing"

// 动物接口
type Animal interface {
	Speak(string) string
}

type Monkey struct {
	Name string
}

func (monkey *Monkey) Speak(word string) string {
	return monkey.Name + "猴子叫" + word
}

type Human struct {
	Name string
}

func (human *Human) Speak(word string) string {
	return human.Name + "人说话" + word
}

func TestOop(t *testing.T) {
	var monkey Animal = &Monkey{"吱吱"}
	t.Log(monkey.Speak("吉吉"))
	var human Animal = &Human{"某"}
	t.Log(human.Speak("测试"))
}
