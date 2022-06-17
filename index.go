package index

import "fmt"


func WrapHandler(handler interface{}) interface{} {

	fmt.Println("Hello from wraphandler")

	return "handler called"
}
