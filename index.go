package index

import "fmt"


func WrapHandler(handler interface{}, nil) string {

	fmt.Println("Hello from wraphandler")

	return handler
}
