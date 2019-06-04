package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str1 = "杨瑞丰斗鱼"
	str, _ := base64.StdEncoding.DecodeString("MEE3OEQyQUJBMTk4NEE1MzgzOEQ4Q0JENEYyN0VDNTJ8NS45MDA")
	fmt.Println(string(str))
	str2 := fmt.Sprintf("str1`s length:%d", len([]rune(str1)))
	println(str2)
}
