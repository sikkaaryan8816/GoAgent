package index

import (
	"fmt"
	"plugin"
)
//import "github.com/aws/aws-lambda-go/lambda"

type LambdaResponse struct {
        PlainText   string
        EncryptedC4 string
        DecryptedC4 string
        EncryptedV  string
        DecryptedV  string
}
type encryptionEngine interface {
        Encrypt(string) string
        Decrypt() (*string, error)
}
func WrapHandler(handler interface{}) interface{} {

	fmt.Println("Hello from wraphandler",*handler)
	handlercall()
	
	return handler
}
func handlercall() (LambdaResponse, error) {
        resp := LambdaResponse{}
	fmt.Println("hello world in golang")
        // Load Cipher plugin
        pluginModule, err := plugin.Open("/opt/cipher.so")
        if err != nil {
                return resp, err
        }
        //Load EncryptCaesar function
        encryptCaesarSymbol, err := pluginModule.Lookup("EncryptCaesar")
        if err != nil {
                return resp, err
        }
        //Load DecryptCaesar function
        decryptCaesarSymbol, err := pluginModule.Lookup("DecryptCaesar")
        if err != nil {
                return resp, err
        }
        //Load VermanCipher variable
        vermanCipherSymbol, err := pluginModule.Lookup("VermanCipher")
        if err != nil {
                return resp, err
        }

        //Cast encryptCaesar symbol to the correct type
        encryptCaesarFunc := encryptCaesarSymbol.(func(int, string) string)
        //Cast encryptCaesar symbol to the correct type
        decryptCaesarFunc := decryptCaesarSymbol.(func(int, string) string)
        //Cast vermanCipher symbol to the correct interface type
        vermanCipherIf := vermanCipherSymbol.(encryptionEngine)

        resp.PlainText = "My name is Abhimanyu kumar."
        resp.EncryptedC4 = encryptCaesarFunc(4, resp.PlainText)
        resp.DecryptedC4 = decryptCaesarFunc(4, resp.EncryptedC4)
        resp.EncryptedV = vermanCipherIf.Encrypt(resp.PlainText)
        decryptedV, err := vermanCipherIf.Decrypt()
        if err != nil {
                return resp, err
        }
        resp.DecryptedV = *decryptedV
        return resp, nil
}
