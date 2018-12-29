
package main
import "fmt"
func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}
func main() {
	array := [5]int{1,2,3,4,5} // 定义并初始化一个数组
	modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容
	fmt.Println("In main(), array values:", array)
	s:=fmt.Sprintf("select tunnel_id,gw_id from t_vpn_tunnel where is_deleted=0 and ipsec_subnets like '%%%s%%'","sunbnet-123456")
	fmt.Println(s)
}
