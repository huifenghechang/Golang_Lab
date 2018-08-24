package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
	"encoding/json"
	"bytes"
	"time"
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
	}else if fc == "transferMarblesBasedOnColor"{
		return m.transferMarblesBasedOnColor(stub,args)
	}else if fc == "delete"{
		return m.delete(stub,args)
	}else if fc == "readMarble"{
		return m.readMarble(stub,args)
	}else if fc == "queryMarblesByOwner"{
		return m.queryMarblesByOwner(stub,args)
	}else if fc == "queryMarbles"{
		return m.queryMarbles(stub,args)
	} else if fc == "getHistoryForMarble" {
		return m.getHistoryForMarble(stub, args)
	} else if fc == "getMarblesByRange" {
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


// 更改Marble的拥有者 ["transferMarble"."marble2","jerry"]
/*
	基本逻辑为“先检验参数合法性、读取参数、从账本中中读取对应值、更改参数、写入账本
*/
func (m *marbleChainCode)transferMarble(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	var marbleName string
	var newOwner string
	var err error
	if len(args) !=2{
		return shim.Error("Incorrect number of Arguments, Excepting 2")
	}

	marbleName = args[0]
	newOwner = args[1]

	marbleAsBytes, err := stub.GetState(marbleName)
	if err != nil{
		return shim.Error("Failed to read "+ marbleName +"from ledger")
	}
	if marbleAsBytes == nil{
		return shim.Error("the marble"+ marbleName+"is not existed!")
	}

	// 将读取的值，进行反序列化成结构体
	marbleObj := marble{}
	err = json.Unmarshal(marbleAsBytes,&marbleObj)
	if err != nil{
		return shim.Error("Failed in the process of json---> struct\n"+err.Error())
	}
	// 更改Marble的所有者
	marbleObj.Owner = newOwner

	marbleAsBytes, err = json.Marshal(marbleObj)
	if err != nil{
		return shim.Error(err.Error())
	}

	err = stub.PutState(marbleName,marbleAsBytes)
	if err != nil{
		return  shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 更改指定颜色的所有大理石的拥有者 ["transferMarblesBasedOnColor","blue","jerry"]
/*
	这里编程的基本逻辑也是：检查参数、读取、更改。

*/
func (m *marbleChainCode)transferMarblesBasedOnColor(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args) != 2{
		return shim.Error("Incorrect number of arguments. Excepting 2")
	}

	color := args[0]
	newOwner := args[1]

	coloredMarbleResultsIterator, err := stub.GetStateByPartialCompositeKey("color~name",[]string{color})
	if err != nil{
		return shim.Error(err.Error())
	}

	defer  coloredMarbleResultsIterator.Close()

	// 遍历每一个marble，更改其拥有者
	var i int
	for i = 0;coloredMarbleResultsIterator.HasNext();i++{
		responseRange, err := coloredMarbleResultsIterator.Next()
		if err != nil{
			return shim.Error(err.Error())
		}

		//  从复合键中分离出color 和 name
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil{
			return shim.Error(err.Error())
		}
		returnColor := compositeKeyParts[0]
		returnMarbleName := compositeKeyParts[1]

		fmt.Printf("- found a marble form index:%s color:%s name %s",objectType, returnColor, returnMarbleName)

		response := m.transferMarble(stub,[]string{returnMarbleName,newOwner})

		if response.Status != shim.OK{
			return shim.Error("Transfer failed:" + response.Message)
		}
	}

	responsePayload := fmt.Sprintf("Transferred %d %s marble to %s", i, color, newOwner)
	fmt.Println("- end transferMarbleBasedOnColor:" + responsePayload)
	return shim.Success(nil)
}

// 返回指定拥有者拥有的所有大理石的信息 ["queryMarble","{\"\selector":{\"owner\":\"Tom\"}}"]
func (m *marbleChainCode)queryMarbles(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	// 根据查询字符串来对账本进行查询
	if len(args) < 1{
		return shim.Error("Incorrect number of arguments.Excepting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}


// 进行富查询！！！
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface,queryString string)([]byte, error)  {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil{
		return nil, err
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext(){
		queryResponse, err := resultsIterator.Next()
		if err != nil{
			return nil, err
		}

		if bArrayMemberAlreadyWritten == true{
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQUeryString queryResult:\n%s\n",buffer.String())

	return buffer.Bytes(),nil


}

// 输入参数["queryMarbleByOwner","tom"]
func (m *marbleChainCode)queryMarblesByOwner(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Excepting 1")
	}

	owner := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\",\"%s\"}}",owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

// 获取历史的marble
func (m *marbleChainCode)getHistoryForMarble(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marbleName := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", marbleName)

	resultsIterator , err := stub.GetHistoryForKey(marbleName)
	if err != nil{
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext(){
		response, err := resultsIterator.Next()
		if err != nil{
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true{
			buffer.WriteString(",")
		}

		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")
		buffer.WriteString(",\"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n",buffer.String())

	return shim.Success(buffer.Bytes())
}

// 返回一定返回内的Marble ["getMarbleByRange","marble1","marble3"]
/*
	基本思路：
		- 检验参数合法性
		- 从账本中返回迭代器
   		- 将迭代器转化为json数组返回之
*/
func (m *marbleChainCode)getMarblesByRange(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args) != 2{
		return shim.Error("Incorrect Number of arguments. Excepting 2")
	}

	fromKey := args[0]
	endKey := args[1]

	resultIterator, err := stub.GetStateByRange(fromKey,endKey)

	if err != nil{
		return shim.Error(err.Error())
	}

	// 务必记住，调用了Iterator迭代器之后，需要用Close方法将其关闭
	defer resultIterator.Close()

	// buffer 是一个包含查询结果数据的json数组
	var  buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultIterator.HasNext(){
		queryResponse, err := resultIterator.Next()
		if err != nil{
			return shim.Error(err.Error())
		}

		// 判断是否加“,”号
		if bArrayMemberAlreadyWritten == true{
			buffer.WriteString(",")
		}

		buffer.WriteString("{\"key\":}")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(",\"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getMarbleByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}


