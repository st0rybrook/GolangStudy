package main

import (
"fmt"
"reflect"
)

type person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func main() {
	t := reflect.TypeOf(person{})
	f1, _ := t.FieldByName("Id")
	fmt.Println(f1.Tag) // f one
	f4, _ := t.FieldByName("FirstName")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t.FieldByName("LastName")
	fmt.Println(f5.Tag) // f four and five
	v := f5.Tag.Get("json")
	fmt.Println(v)
	v, ok := f5.Tag.Lookup("json")
	fmt.Println("?,?", v, ok)

}
