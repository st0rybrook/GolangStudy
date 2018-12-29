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
				<span>
					<div>DIV5</div>
				</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div:only-child").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}
