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
