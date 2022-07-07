package index

import (
	"fmt"
	"context"
	"encoding/json"
	"errors"
	"reflect"
)
//import "github.com/aws/aws-lambda-go/lambda"

var (
	// CurrentContext is the last create lambda context object.
	CurrentContext context.Context
)

func WrapHandler(handler interface{}) interface{}  {

	fmt.Println("Hello from wraphandler= %v ",handler)
	
	
	result, err := callHandler(ctx, msg, handler)
	if err == nil {
		fmt.Println("error")
	}
	return result
}

func callHandler(ctx context.Context, msg json.RawMessage, handler interface{}) (interface{}, error) {
	
	handlerType := reflect.TypeOf(handler)

	args := []reflect.Value{}

	if handlerType.NumIn() == 1 {
		// When there is only one argument, argument is either the event payload, or the context.
		contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
		firstArgType := handlerType.In(0)
		if firstArgType.Implements(contextType) {
			args = []reflect.Value{reflect.ValueOf(ctx)}
		} else {
			args = []reflect.Value{ev.Elem()}

		}
	} else if handlerType.NumIn() == 2 {
		// Or when there are two arguments, context is always first, followed by event payload.
		args = []reflect.Value{reflect.ValueOf(ctx), ev.Elem()}
	}

	handlerValue := reflect.ValueOf(handler)
	output := handlerValue.Call(args)

	var response interface{}
	var errResponse error

	if len(output) > 0 {
		// If there are any output values, the last should always be an error
		val := output[len(output)-1].Interface()
		if errVal, ok := val.(error); ok {
			errResponse = errVal
		}
	}

	if len(output) > 1 {
		// If there is more than one output value, the first should be the response payload.
		response = output[0].Interface()
	}

	return response, errResponse
}

