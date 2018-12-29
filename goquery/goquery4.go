package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	html := `<body>

				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[class]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	dom.Find("div[class=name]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	//选择器	说明
	//Find(“div[lang]“)	筛选含有lang属性的div元素
	//Find(“div[lang=zh]“)	筛选lang属性为zh的div元素
	//Find(“div[lang!=zh]“)	筛选lang属性不等于zh的div元素
	//Find(“div[lang¦=zh]“)	筛选lang属性为zh或者zh-开头的div元素
	//Find(“div[lang*=zh]“)	筛选lang属性包含zh这个字符串的div元素
	//Find(“div[lang~=zh]“)	筛选lang属性包含zh这个单词的div元素，单词以空格分开的
	//Find(“div[lang$=zh]“)	筛选lang属性以zh结尾的div元素，区分大小写
	//Find(“div[lang^=zh]“)	筛选lang属性以zh开头的div元素，区分大小写
}
