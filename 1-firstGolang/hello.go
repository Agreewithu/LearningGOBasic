package main //当前函数的包名
import (
	"fmt"
	"time"
)

//main函数
func main() {
	fmt.Println("Hello, Go world~")
	//睡眠1s
	time.Sleep(1 * time.Second)
}
