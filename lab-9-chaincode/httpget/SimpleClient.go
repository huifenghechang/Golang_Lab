package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"net/http"
)

type SimpleClient struct {
}

// Init、Invoke 函数为chainCode接口的两个子函数
func (Client *SimpleClient) Init(stub shim.ChaincodeStubInterface) pb.Response {
	/*args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments.Excepting a key and a value")
	}

	err := stub.PutState(args[0],[]byte(args[1]))
	if err != nil{
		return shim.Error(fmt.Sprintf("Failed to create SimpleClient: %s",err))
	}*/
	return shim.Success(nil)

}

func (Client *SimpleClient) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "getUrl" {
		result = getUrl()
	} else {
		result = postUrl()
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("the args is %s", args)
	return shim.Success([]byte(result))

}

func getUrl() string {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// 错误处理
		// shim.Error(err.Error())
		return "get Url Failed"
	}

	defer resp.Body.Close() //关闭链接

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v= %v\n", k, v)
	}

	fmt.Printf("resp status %s, statusCode %d \n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto %s \n", resp.Proto) // proto 协议
	fmt.Printf("resp content length %d \n", resp.ContentLength)

	return "get Url Succeed!"
}

func postUrl() string {
	fmt.Printf("Hello-World!This is BlockChain World!")
	return "Post Succeed!"
}

func main() {
	if err := shim.Start(new(SimpleClient)); err != nil {
		fmt.Printf("Dear sxl,Error starting SimpleClient chaincode %s", err)
	}

}
