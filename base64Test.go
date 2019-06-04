package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	str,_:=base64.StdEncoding.DecodeString("MEE3OEQyQUJBMTk4NEE1MzgzOEQ4Q0JENEYyN0VDNTJ8NS45MDA")
	fmt.Println(string(str))
}
