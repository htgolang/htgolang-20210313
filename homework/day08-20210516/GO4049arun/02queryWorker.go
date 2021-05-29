package main
import (
	"fmt"
	"sync"
)
func main() {
	// pool
	// 需要资源 从池子里面获取,当池子中没有会进行创建并返回,有直接复用,用完后归还,资源的重复利用
	// 例如:对象创建时耗费资源,网络连接池,数据库连接池
	// 创建资源 => func 创建资源方式
	// 获取资源
	// 放入资源
	objPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("new")
			return "123456789"
		},
	}
	obj := objPool.Get()    		//new
	fmt.Println(obj)				//123456789
	obj2 := objPool.Get()			//new
	fmt.Println(obj2)				//123456789

	objPool.Put(obj)
	obj3 := objPool.Get()
	fmt.Println(obj3)				//123456789
}