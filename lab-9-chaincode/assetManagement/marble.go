package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
	"encoding/json"
)

/*
本例中，主要是学会在ChainCode中定义一项资产，并围绕该资产提供创建、查询、转移所有权等操作。
*/

type marbleChainCode struct {

}

// 此处用到了Struct Tag 相关语法
// 为每一个成员变量设定了标签（如json："docType"），用于指定将结构体序列化成特定格式（如JSON）时该字段的键的名称。
// 定义了资产类型marble
type marble struct {
	ObjectType string `json:"docType"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}

// 定义主函数，用于向peer节点注册chainCode
func main(){
	err := shim.Start(new(marbleChainCode))
	if err != nil{
		fmt.Printf("Error starting Simple Chaincode %s",err)
	}
}

// 实现ChainCode接口
func (m *marbleChainCode)Init(stub shim.ChaincodeStubInterface) pb.Response{
	// 不做任何处理
	return shim.Success(nil)
}

func (m *marbleChainCode)Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	// 将不同的请求，路由至不同的函数
	fc, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke is running")

	// 针对不同的函数名，定位到不同的处理逻辑上去
	if fc == "initMarble"{
		return m.initMarble(stub,args)
	}else if fc == "transferMarble"{
		return m.transferMarble(stub,args)
	}else if fc == "transferMarbleBaseOnColor"{
		return m.transferMarbleBaseOnColor(stub,args)
	}else if fc == "delete"{
		return m.delete(stub,args)
	}else if fc == "readMarble"{
		return m.readMarble(stub,args)
	}else if fc == "queryMarbleByOwer"{
		return m.queryMarbleByOwer(stub,args)
	}else if fc == "queryMarbles"{
		return m.queryMarbles(stub,args)
	} else if function == "getHistoryForMarble" {
		return m.getHistoryForMarble(stub, args)
	} else if function == "getMarblesByRange" {
		return m.getMarblesByRange(stub, args)
	}
	return shim.Error("Received Unknown function invocation")
}

// 初始化Marble

/* 基本思路
   1.检验传入参数合法性 {"Init":["marble1","red","88","Tom"]}
   2.检查创建的Marble是否存在
   3.创建结构体，序列化为json
   4.将数据存入账本
   5.创在复合键。以方便查询
*/
func (m *marbleChainCode)initMarble(stub shim.ChaincodeStubInterface,args []string) pb.Response{

	// 检验传入参数的合法性
	if len(args) != 4{
		return shim.Error("Incorrect number of args. Excepting 4")
	}
	fmt.Println("- start init marble")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}

	// 解析传入的参数
	marbleName := strings.ToLower(args[0])
	color := strings.ToLower(args[1])
	size, err := strconv.Atoi(args[2])
	ower := strings.ToLower(args[3])

	if err != nil{
		return shim.Error("3rd arguments must be a numeric string")
	}

	// 检查创建的Marble是否存在于账本中
	marbleAsBytes, err := stub.GetState(marbleName)
	if err != nil{
		return shim.Error("Failed to get marble:"+ err.Error())
	}
	if marbleAsBytes != nil{
		return shim.Error("This marble is exits:" + marbleName)
	}

	//创建Marble结构体
	objectType := "marble"
	marble := &marble{objectType,marbleName,color,size,ower}
	marbleJSONBytes, err := json.Marshal(marble)
	if err != nil{
		return shim.Error(err.Error())
	}

	// 将新创建的结构体存于账本中
	err = stub.PutState(marbleName,marbleJSONBytes)
	if err != nil{
		return shim.Error(err.Error())
	}

	// 创建复合键，来辅助查询
	// 创建该复合键，是为了对某一特定颜色的大理石进行范围查找
	// 这里，复合键的意义是将一部分属性也构造为了索引的一部分
	indexName := "color-name"
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName,[]string{marble.Color,marble.Name})
	if err != nil{
		return shim.Error(err.Error())
	}
	// 以复合键为键、0x00为值，将复合键记录到账本中
	value := []byte{0x00}
	stub.PutState(colorNameIndexKey,value)
	fmt.Println(" -end init marble")
	return shim.Success(nil)
}


// readMarble, 用于读取Marble的值，并返回
// 输入参数形如：["reaMarble","marble1"]
func (m *marbleChainCode)readMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Excepting name of the marble to query")
	}

	marbleName := args[0]
	marbleBytes, err := stub.GetState(marbleName)

	if err != nil{
		return shim.Error("Failed to get"+ marbleName+ err.Error())
	}
	if marbleBytes == nil{
		return shim.Error(marbleName+"not exist ! ")
	}

	// 将读取的数据字符串化，并输出
	marbleValue := string(marbleBytes)
	fmt.Println(marbleValue)
	return shim.Success(marbleBytes)
}


// 删除marble
// 传入的参数为 ["delete","marble1"]
// 删除之前，需要先从账本中读取，确保该marble存在。
func (m *marbleChainCode)delete(stub shim.ChaincodeStubInterface,args []string) pb.Response  {
	var marbleName string
	var err error
	var marbleJson marble

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.Excepting name of marble to delete")
	}

	marbleName = args[0]

	valAsbytes, err := stub.GetState(marbleName)
	if err != nil{
		jsonResp := "{\"Error\":\"Failed to get state for " +marbleName+ "\"}"
		return shim.Error(jsonResp)
	}
	if valAsbytes == nil{
		jsonResp :="{\"Error\":\"Marble does not exist !"+ marbleName+ "\"}"
		return shim.Error(jsonResp)
	}

	// 将json转换为结构体
	err = json.Unmarshal([]byte(valAsbytes),&marbleJson)
	if err != nil{
		jsonResp := "{\"Error\":\"Failed to decode JSON of: " + marbleName + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(marbleName)
	if err != nil{
		return shim.Error("Failed to delete state!" + err.Error())
	}
	return shim.Success(nil)

}