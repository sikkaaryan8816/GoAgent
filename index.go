package index

import "fmt"


func WrapHandler(handler interface{}, a string)) lambda.Handler {

	fmt.Println("Hello from wraphandler",a)

	return &handler
}
