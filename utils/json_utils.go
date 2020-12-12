package utils

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type JsonUtils struct {
}

func ObjToJsonStr(obj interface{}) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary //号称全宇宙最快的json解析包
	toString, err2 := json.MarshalToString(obj)
	err := err2
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	return toString
}
