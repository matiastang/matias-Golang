package main

import (
	"fmt"
	"test/start"
)

func init()  {
	fmt.Println("main init")
	start.Hello()
}

func main()  {
	fmt.Println("main main")
	start.Hello()
}