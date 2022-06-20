package index

import "fmt"


func WrapHandler(handler interface{}, nil) lambda.Handler {

	fmt.Println("Hello from wraphandler",a)

	return handler
}
