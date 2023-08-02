package main

import (
	"fmt"
	"tutorial-golang/helper"
)

func main() {
	helper.SayHello("Thomas")
	// helper.sayGoodbye("Thomas") //error
	fmt.Println(helper.Application)
	fmt.Println(helper.version)
}