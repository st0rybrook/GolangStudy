package main

import "fmt"
import (
	"html/template"
	"os"
)

//func Hello(response http.ResponseWriter, request *http.Request) {
//	type person struct {
//		Id      int
//		Name    string
//		Country string
//	}
//
//	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "China"}
//
//	tmpl, err := template.ParseFiles("D:\\GoPro\\src\\github.com\\golang-china\\GolangStudy\\bootstrap\\user.tpl")
//	if err != nil {
//		fmt.Println("Error happened..")
//	}
//	tmpl.Execute(response, liumiaocn)
//}
//
//func main() {
//	http.HandleFunc("/", Hello)
//	http.ListenAndServe(":8080", nil)
//}

func main() {
	type person struct {
		Id      int
		Name    string
		Country string
	}

	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "China"}

	tmpl, err := template.ParseFiles("D:\\GoPro\\src\\github.com\\golang-china\\GolangStudy\\bootstrap\\user.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(os.Stdout, liumiaocn)
}