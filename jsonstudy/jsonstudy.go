package main

import (
	"os"
	"encoding/json"
	"reflect"
	"io/ioutil"
	"fmt"
)

func main()  {
	fileName := "D:\\GoPro\\src\\github.com\\golang-china\\GolangStudy\\jsonstudy\\vendor.json"
	f,err:= os.Open(fileName)
	if err!=nil{
		panic("open file failed")
	}
	v,_:= ioutil.ReadAll(f)
	vf:=reflect.TypeOf(v)
	fmt.Println(vf)
	jsonMap := map[string]interface{}{}
	json.Unmarshal(v, &jsonMap)
	fmt.Println(jsonMap)
	configMap := map[string]interface{}{}
	for k, v := range jsonMap {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			for k1, v1 := range v.(map[string]interface{}) {
				if reflect.TypeOf(v1).Kind() == reflect.Map {
					for k2, v2 := range v1.(map[string]interface{}) {
						configMap[k2] = v2
					}
				} else {
					configMap[k1] = v1
				}
			}
		} else {
			configMap[k] = v
		}
	}
}
