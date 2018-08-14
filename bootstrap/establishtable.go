package main

//传入数据构建html表格元素
import (
	"html/template"
	"log"
	"reflect"
	"strconv"
	"bytes"
	"fmt"
)

//定义HTML模板,邮件只支持原生HTMl
const tpl = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8"> 
	<title>显示结果</title>
	</head>
<body>

<table border="1px" bordercolor="#000000" cellspacing="0px" style="border-collapse:collapse" align="center">
	<caption{{ .TableName}}</caption>
	<thead>
		<tr>
			{{range .HeadNameList}}<th align="center">{{ . }}</th>{{end}}
			
		</tr>
	</thead>
	<tbody>

		{{range  $i, $v :=.Data}}
		<tr>
			{{range $ii, $vv := $v }}<td bgcolor={{confirmGridColor $vv $ii $.HeadNameList $.ParaData}} align="center"> {{$vv}}</td>{{end}}
		</tr>
		{{end}}
	</tbody>
</table>

</body>
</html>`

//首字母大写，否则无法反射类型
type HostInfo struct {
	Host     string
	EipCount int
	NsCount  int
	Pps      int
	Bps      int
	Cpu      int
	Score    int
	Zone     int
	Region   int
}
type MailHead map[string]int

//表格每行信息
type HostInfoList []string

//详细表格信息
type TableInfos struct {
	TableName    string         //表格名称
	HeadNameList []string       //表头信息
	Data         []HostInfoList //单元格信息
	ParaData     map[string]int //比较数据
}

//颜色常量
const (
	Red   = "#FF0000"
	White = "#FFFFFF"
)

//比较数据大小，确定单元格颜色，大于阈值设置红色
func confirmGridColor(value string, index int, HeadNameList []string, paraData MailHead) string {
	if index == 0 {
		return White
	}
	va, err := strconv.Atoi(value)
	if err != nil {
		panic("can not convert value to int ")
	}
	if _, ok := paraData[HeadNameList[index]]; ok {
		if va > paraData[HeadNameList[index]] {
			fmt.Println(va, "  ", paraData[HeadNameList[index]])
			return Red

		} else {
			return White
		}

	} else {
		panic("can not find data in MailHead")
	}
	return White
}

//反射获取表头的信息，用于打印table表头
func getHeadList(o interface{}) []string {
	var result = []string{}
	t := reflect.TypeOf(o)
	//fmt.Println("Fileds:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		result = append(result, f.Name)

	}
	return result
}

//反射遍历Hostinfo结构体, html模板无法遍历struct, 返回list
func getHostInfo(hostInfo []HostInfo) []HostInfoList {
	var result = []HostInfoList{}
	for _, hi := range hostInfo {
		var li = HostInfoList{}
		t := reflect.TypeOf(hi)
		v := reflect.ValueOf(hi)
		for i := 0; i < t.NumField(); i++ {
			//f := t.Field(i)
			val := v.Field(i).Interface()
			if reflect.TypeOf(val).Name() == "int" {
				val = strconv.Itoa(val.(int))
			}
			li = append(li, val.(string))
		}
		result = append(result, li)
	}
	return result
}

//解析主函数
func parseMailStr(mailHead MailHead, data []HostInfo) string {

	var table TableInfos
	table.Data = getHostInfo(data)
	table.ParaData = mailHead
	table.TableName = "表一"
	table.HeadNameList = getHeadList(data[0])

	//数据不能为空
	if len(table.HeadNameList) == 0 {
		panic("error: host infomation is empty!!!")
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	//解析HTML模板及函数模板
	//t, err := template.New("report").Funcs(template.FuncMap{"confirmGridColor": confirmGridColor}).Parse(tpl)
	t, err := template.New("index.html").Funcs(template.FuncMap{"confirmGridColor": confirmGridColor}).ParseFiles("D:\\GoPro\\src\\github.com\\golang-china\\GolangStudy\\bootstrap\\index.html", "D:\\GoPro\\src\\github.com\\golang-china\\GolangStudy\\bootstrap\\table.html")
	check(err)
	//var str string
	buf := bytes.NewBufferString("")

	//结构类似的多个table列表在模板中循环遍历，要改一下传入的数据
	err = t.Execute(buf, []TableInfos{table, table})
	check(err)
	return buf.String()
}

//测试一下
func main() {
	hostInfo := HostInfo{
		"192.168.1.1",
		1,
		2,
		3,
		4,
		5,
		6,
		6,
		7,
	}
	hostInfo2 := HostInfo{
		"192.168.1.2",
		1,
		2,
		3,
		0,
		1,
		6,
		1,
		7,
	}
	mp := map[string]int{
		"EipCount": 1,
		"NsCount":  5,
		"Pps":      5,
		"Bps":      5,
		"Cpu":      5,
		"Score":    5,
		"Zone":     5,
		"Region":   5,
	}
	fmt.Println(parseMailStr(mp, []HostInfo{hostInfo, hostInfo2}))
}
