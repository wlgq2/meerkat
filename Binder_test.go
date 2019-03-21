package meerkat


import  "testing"
import  "reflect"

type TestBindData struct {
	Num1 int   `json:"num1" form:"num1" query:"num1"`
	Num2 uint  `json:"num2" form:"num2" query:"num2"`
	Flag bool   `json:"flag" form:"flag" query:"flag"`
	Str string `json:"str" form:"str" query:"str"`
}

func TestBindForm(t *testing.T) {
	data :=TestData{}
	strs := make(map[string][]string)
	strs["num1"] = []string{"-1"}
	strs["num2"] = []string{"1024"}
	strs["str"] = []string{"test"}
	strs["flag"] = []string{"true"}
	bindData(&data,strs,"form")
	if data.Num1 != -1 {
		t.Error("test bind form error :",reflect.TypeOf(data.Num1))
	}
	if data.Num2 != 1024 {
		t.Error("test bind form error :",reflect.TypeOf(data.Num2))
	}
	if data.Str != "test" {
		t.Error("test bind form error :",reflect.TypeOf(data.Str))
	}
	if !data.Flag {
		t.Error("test bind form error :",reflect.TypeOf(data.Flag))
	}
	t.Log("success~")
}
