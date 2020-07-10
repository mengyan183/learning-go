package polymorphic

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type JavaProgrammer struct {
}

// 重写
func (p *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.Println(\"HelloWorld\")"
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"HelloWorld\")"
}

func TestPolymorphic(t *testing.T) {
	var p Programmer
	p = &JavaProgrammer{}
	t.Logf("%T ; %s \n", p, p.WriteHelloWorld())
	p = &GoProgrammer{}
	t.Logf("%T ; %s \n", p, p.WriteHelloWorld())
}

func Polymorphic(p Programmer)  {
	p.WriteHelloWorld()
}


