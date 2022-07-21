package index

import (
	"fmt"
	"context"
	"encoding/json"
	//"errors"
	"reflect"
)
//import "github.com/aws/aws-lambda-go/lambda"

var (
	// CurrentContext is the last create lambda context object.
	CurrentContext context.Context
)

func WrapHandler(handler interface{}) interface{}  {

	fmt.Println("Hello from wraphandler= %v ",handler)
	coldStart := true

	
	return func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
		//nolint
		ctx = context.WithValue(ctx, "cold_start", coldStart)
		/*for _, listener := range listeners {
			ctx = listener.HandlerStarted(ctx, msg)
		}*/
		//UDPConnection()
		aws_request_id := "default_aws_request_id"
		function_name := "defult"
		url_path := context_obj.function_name
	
		bt := start_business_transaction(url_path,"") 
	
		handle, handler_name, mpackage := "test","Test","TEST"
	
		fqmmethodentry := "handle.main.Test"
	
		method_entry(bt,fqmmethodentry)
		
		CurrentContext = ctx
		result, err := callHandler(ctx, msg, handler)
		/*for _, listener := range listeners {
			listener.HandlerFinished(ctx, err)
		}*/
		method_exit(bt,fqmmethodentry,200)
	
		end_business_transaction(bt)
		coldStart = false
		CurrentContext = nil
		return result, err
	}
}

func callHandler(ctx context.Context, msg json.RawMessage, handler interface{}) (interface{}, error) {
	ev, err := unmarshalEventForHandler(msg, handler)
	if err != nil {
		return nil, err
	}
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
func unmarshalEventForHandler(ev json.RawMessage, handler interface{}) (reflect.Value, error) {
	handlerType := reflect.TypeOf(handler)
	if handlerType.NumIn() == 0 {
		return reflect.ValueOf(nil), nil
	}

	messageType := handlerType.In(handlerType.NumIn() - 1)
	contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
	firstArgType := handlerType.In(0)

	if handlerType.NumIn() == 1 && firstArgType.Implements(contextType) {
		return reflect.ValueOf(nil), nil
	}

	newMessage := reflect.New(messageType)
	err := json.Unmarshal(ev, newMessage.Interface())
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return newMessage, err
}

func start_business_transaction(bt_name ,correlation_header string){
		
		UDPConnection()
		message := create_start_transaction_message(context,bt_name,correlation_header)
		udp_connection.send(massage,"start_fp")
		bt := generate_bt()
		
		return bt
}
func end_business_transaction(bt int64){
		message := create_end_transaction_message(context,bt,200)
		udp_connection.send(massage)
		rc := 0
		return rc
	}


