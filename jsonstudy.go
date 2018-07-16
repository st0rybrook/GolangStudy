package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func Println(book *Book) {
	fmt.Println(book.Title)
	fmt.Println(book.Authors)
	fmt.Println(book.Publisher)
	fmt.Println(book.IsPublished)
	fmt.Println(book.Price)

}
func main() {
	var s = []string{"tang", "shen"}

	gobook := Book{
		"Go语言编程",
		s,
		"ituring.com.cn",
		true,
		9.0}
	b, err := json.Marshal(gobook)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(b)
	Println(&gobook)
	c := []byte(`{
 "Title": "Go语言编程",
 "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
 "XuDaoli"],
 "Publisher": "ituring.com.cn",
 "IsPublished": true,
 "Price": 9.99,
 "Sales": 1000000
}`)
	var r interface{}
	err1 := json.Unmarshal(c, &r)
	if err1 != nil {
		fmt.Println("error")
	}
	gobook1, ok := r.(map[string]interface{})
	if ok {
		for k, v := range gobook1 {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}
