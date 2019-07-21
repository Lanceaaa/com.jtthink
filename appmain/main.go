package main

import (
	"fmt"
    _ "com.jtthink/services"
	. "com.jtthink/core"
)

func main()  {
	fmt.Println(GetService().Get(123))
}