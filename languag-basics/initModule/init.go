package init

import "fmt"

func init() {
	fmt.Println("Init was called when package has been imported")
}

func Hello() string {
	fmt.Println("Hello was called")
	return "Hello"
}
