package main

import (
	"nockberge/service"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	service.Start()
	service.Stop()
}
