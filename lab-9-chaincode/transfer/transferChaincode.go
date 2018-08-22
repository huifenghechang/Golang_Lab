package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"fmt"
)

type SimpleChainCode struct {

}

func (s *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response{
	args := stub.GetStringArgs()
	if len(args) != 4 {
		return shim.Error("Incorrect of number of argument. Excepting 4")
	}
	//读取参数
	A := args[0]
	Aval, err := strconv.Atoi(args[1])
	if err != nil{
		return shim.Error("Excepting an integer value for asset holding")
	}

	B := args[2]
	Bval, err := strconv.Atoi(args[3])
	if err != nil{
		return shim.Error("Excepting an integer value for asset holding")
	}

	// 将对应的值写入到数据库中
	err = stub.PutState(A,[]byte(Aval))
	if err != nil{
		return shim.Error(err.Error())
	}

	err = stub.PutState(B,[]byte(Bval))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (s *SimpleChainCode)Invoke(stub shim.ChaincodeStubInterface) pb.Response  {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("transferChaincode Invoke running")

	// 根据获取的函数名，对需要调用的函数进行路由
	if fn == "query" {
		return s.query(stub, args)
	}else if fn == "invoke" {
		return s.invoke(stub, args)
	}else if fn == "delete" {
		return s.delete(stub, args)
	}

	return shim.Error("Invalid func name. It must be \"query\" \"invoke\" or \"delete\"")

}

func (s *SimpleChainCode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response  {

	// 编程心得：在函数的开头，定义需要用到的变量
	var A string
	var err error

	// 检查args的合法性
	if len(args) != 1 {
		return shim.Error("Error number of param . Excepting Account name")
	}

	 A = args[0]
	 Avalbytes, err := stub.GetState(A)
	 if err != nil{
	 	Resp := "Failed to get state for " + A
	 	return shim.Error(err.Error()+"\n"+Resp)
	 }
	 if Avalbytes == nil{
	 	Resp := "Nil amount for " + A
	 	return shim.Error(Resp)
	 }

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Println(jsonResp)
	return shim.Success(Avalbytes)
}


func (s *SimpleChainCode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string
	var Aval, Bval int
	var x int
	var err error

	if len(args) != 3{
		return shim.Error("Excepting three params like \"{\"a\",\"b\",\"10\"}\"")
	}

	A = args[0]
	B = args[1]

	// 从数据库中查询A对应的参数值
	Avalbytes, err := stub.GetState(A)
	if err != nil{
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil{
		return shim.Error("Entity "+ A +"not found!")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	// 从数据库中查询B对应的数值
	Bvalbytes, err := stub.GetState(B)
	if err!= nil{
		return shim.Error("Failed to get B")
	}
	if Bvalbytes == nil {
		shim.Error("Entity"+ B +"not found")

	}

	// 进行转账操作
	// {"a"，"b","10"},表示将a用户转账给b用户10元
	x, err = strconv.Atoi(args[2])
	if err != nil{
		return shim.Error("Failed to exec Transaction.the third params required an integer")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))
	Bval, _ = strconv.Atoi(string(Bvalbytes))
	Aval = Aval - x
	Bval = Bval + x
	fmt.Printf("Aval = %d, Bval = %d\n",Aval,Bval)
	
	//将新的世界状态写入到账本中
	err = stub.PutState(A,[]byte(strconv.Itoa(Aval)))
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(B,[]byte(strconv.Itoa(Bval)))
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


/* delete 函数，参入的参数为{"delete","a"}的形式。
所以，只需要获取对应的参数，然后调用stub接口，对账本进行操作即可
*/
func (s *SimpleChainCode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response  {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Excepting 1")
	}
	
	A := args[0]
	err := stub.DelState(A)
	if err != nil{
		shim.Error("Failed to delete state")
	}
	return shim.Success(nil)
}


// shim.Start()的功能：向peer节点注册chainCode
func main()  {
	err := shim.Start(new(SimpleChainCode))
	if err != nil{
		fmt.Printf("Error starting SimpleChainCode:%s",err)
	}


}