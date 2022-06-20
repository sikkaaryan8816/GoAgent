package index

import "fmt"
import "github.com/aws/aws-lambda-go/lambda"

func WrapHandler(handler interface{}) interface{} {

	fmt.Println("Hello from wraphandler")

	return handler
}
