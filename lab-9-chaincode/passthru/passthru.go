package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strings"
)

type PassthruChaincode struct {

}

//将接收的参数，转换为ChainCode需要用到的参数
func toChainCodeArgs(args ...string)[][]byte{
	// [][]byte表示的是二维数组，本段代码是将多个[]byte数组合并成一个[]byte
	bargs := make([][]byte,len(args))
	for i,arg := range args{
		bargs[i] = []byte(arg)
	}
	return bargs
}

// 初始化操作
func (p *PassthruChaincode)Init(stub shim.ChaincodeStubInterface) pb.Response  {
	function,_ := stub.GetFunctionAndParameters()
	//strings.Index 的作用： strings.Index(s string, substr string) int
	// 检验function是否含有“error”,返回子串前面的字符的个数
	if strings.Index(function,"error") >= 0{
		return shim.Error(function)
	}
	return shim.Success([]byte(function))
}


func (p *PassthruChaincode)iq(stub shim.ChaincodeStubInterface, function string,args []string)pb.Response{
	if function == ""{
		return shim.Error("ChainCode ID not provided")
	}
	chainCodeID := function

	return stub.InvokeChaincode(chainCodeID,toChainCodeArgs(args...),"")
}


func (p *PassthruChaincode)Invoke(stub shim.ChaincodeStubInterface) pb.Response  {
	function, args := stub.GetFunctionAndParameters()
	return p.iq(stub,function,args)
}


func main()  {
	err := shim.Start(new(PassthruChaincode))
	if err !=nil{
		fmt.Printf(err.Error())
	}

}

