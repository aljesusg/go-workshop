package main

import (
	"fmt"

	"github.com/aljesusg/go-workshop/chapter02_packages/hello_package"
)

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := hello_package.HelloWorld("go developer")
	fmt.Println(hi)
}