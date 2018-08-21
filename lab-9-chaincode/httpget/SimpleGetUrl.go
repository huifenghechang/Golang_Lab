package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"net/http"
)

type SimpleGet struct {
}

// Init函数，主要是对一些数据进行初始化，在智能合约初始化以及升级的时候，会调用该函数。
func (client *SimpleGet) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetStringArgs()
	// 检验传入的参数是否合法
	if len(args) < 2 {
		return shim.Error("Wrong Args !!!")
	}
	// 若合法，则进行相关初始化操作
	for pama, _ := range args {
		fmt.Println(pama)
	}
	fmt.Println(args)
	// 返回成功初始化响应
	return shim.Success(nil)
}

// Invoke函数，主要是用来对事务进行操作。发起交易，则调用智能合约的该函数。
func (client *SimpleGet) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error

	if fn == "getRequest" {
		result, err = getRequest(stub, args)
	} else if fn == "postRequest" {
		result, err = postRequest(stub, args)
	} else {
		sayHello()
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

// 发起http请求，使用get方法获取数据。再将所得到的数据，保存到区块链中。
//  输入的参数为传入的url的值！
func getRequest(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	//因为调用不同函数，传入的参数要求是不一样的，所以，在每个子函数中，需要各自检验参数是否合法
	if len(args) != 1 {
		return "Error Args", fmt.Errorf("Incorrect arguments. Excepting a key")
	}
	url := args[0]

	// 发起http请求
	resp, err := http.Get(url)
	if err != nil {
		return "Error Get Request", fmt.Errorf("Failed to Create Get Request, %s", err.Error())
	}

	defer resp.Body.Close() //关闭链接

	// 对返回的参数，进行验证与存储
	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v", k, v)
	}

	// 可以使用stub.putState()、stub.getStub()等进行读写操作
	err = stub.PutState("suqiancheng", []byte("DaBao"))
	if err != nil {
		return "failed put to blockChain", fmt.Errorf("Failed")
	}

	//打印获取的参数
	fmt.Printf("resp status %s, statusCode %d \n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto %s \n", resp.Proto) // proto 协议
	fmt.Printf("resp content length %d \n", resp.ContentLength)

	buf := bytes.NewBuffer(make([]byte, 0, 512))
	buf.ReadFrom(resp.Body)
	fmt.Println(string(buf.Bytes()))

	return "ALl Right!", nil

}

// 发起post请求，等待服务器响应。最后将处理好的数据保存到区块链中。
func postRequest(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	fmt.Println(args)
	// 可以使用stub.putState()、stub.getStub()等进行读写操作
	err := stub.PutState("suqiancheng", []byte("DaBao"))
	if err != nil {
		return "failed put to blockChain", fmt.Errorf("Failed")
	}

	return "Post not writing", nil

}

// 针对任何其他不存在的函数名字，调用该函数
func sayHello() {
	fmt.Println("Hello World~ \n Welcome to blockChain World!")

}
