package index

import "fmt"


func WrapHandler(handler interface{}) lambda.Handler {

	fmt.Println("Hello from wraphandler",a)

	return handler
}
