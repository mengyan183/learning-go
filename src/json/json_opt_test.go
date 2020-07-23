package json_opt

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

//采用内置的json解析工具
func TestJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*e)
	s, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(s))
}

func TestFastJson(t *testing.T) {
	e := new(Employee)
	//err := easyjson.Unmarshal([]byte(jsonStr), e)
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*e)
	//s, err := easyjson.Marshal(e)
	s, err := e.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(s))
}

func Benchmark_MarshalEasyJSON(b *testing.B) {
	b.StartTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Fatal(err)
		}
		b.Log(*e)
		s, err := e.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
		b.Log(string(s))
	}
}

func Benchmark_JSON(b *testing.B) {
	b.StartTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Fatal(err)
		}
		b.Log(*e)
		s, err := json.Marshal(e)
		if err != nil {
			b.Fatal(err)
		}
		b.Log(string(s))
	}

}
