package struct_test

import "testing"

type Person struct {
	Name string
	Age  int
	Degree
}
type Degree struct {
	Grade int
	Name  string
}

func TestStruct(t *testing.T) {
	p := Person{
		Name: "guoxing",
		Age:  18,
		Degree: Degree{
			Grade: 9,
			Name:  "初中",
		},
	}
	t.Logf("%p", &p)
	t.Logf("%p", &p.Name)
	t.Logf("%p", &p.Age)
	t.Logf("%p", &p.Degree)
	t.Logf("%p", &p.Degree.Grade)
	t.Logf("%p", &p.Degree.Name)
}
