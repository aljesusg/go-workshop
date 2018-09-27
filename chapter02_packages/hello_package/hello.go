package hello_package

import "fmt"

// HelloWorld func say hi
func HelloWorld(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
