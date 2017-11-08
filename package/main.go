package main

import (
	"fmt"
	// 需在`GOROOT`或者`GOPATH`的`src`目录下创建
	// 路径为`package/math`的文件夹。然后将`math.go`拷贝进去
	"package/math"
)

// godoc -http=":6060"
// 可在浏览器查看本机安装的所有包

func main() {
	xs:=[]float64{1,2,3,4}
	avg:=math.Average(xs)
	fmt.Println(avg)
}