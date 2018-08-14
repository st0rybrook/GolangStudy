package main

import (

	"runtime"
	"reflect"
	"strings"
	"fmt"
)



func main() {
	fmt.Println(GetTaskFuncName(tang))
}
func tang() int {
	i := 1
	return i
}
func GetTaskFuncName(taskHandler interface{}) string {
	funcInfo := runtime.FuncForPC(reflect.ValueOf(taskHandler).Pointer()).Name()
	return strings.Split(funcInfo, ".")[1]
}