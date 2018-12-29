package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[lang=zh]+p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("div[lang=zh]~p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("div:contains(DIV2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find(":empty").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
