//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("return:", a()) // 打印结果为 return: 0
//}
//
//func a() int {
//	var i int
//
//	defer func() {
//		i++
//		fmt.Println("defer2:", i) // 打印结果为 defer: 2
//	}()
//	defer func() {
//		i++
//		fmt.Println("defer1:", i) // 打印结果为 defer: 1
//	}()
//	return i
//}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("return:", b()) // 打印结果为 return: 2
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}
/*
先来假设出结论，帮助大家理解

多个defer的执行顺序为“后进先出”；

defer、return、返回值三者的执行逻辑应该是：return最先执行，
return负责将结果写入返回值中；
接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

如何解释两种结果的不同：

上面两段代码的返回结果之所以不同，其实从上面第2条结论很好理解。
a()int 函数的返回值没有被提前声明，其值来自于其他变量的赋值，
而defer中修改的也是其他变量，而非返回值本身，因此函数退出时返回值并没有被改变。

b()(i int) 函数的返回值被提前声明，也就意味着defer中是可以调用到真实返回值的，
因此defer在return赋值返回值 i 之后，再一次地修改了 i 的值，
终函数退出后的返回值才会是defer修改过的值。
*/