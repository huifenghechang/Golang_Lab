package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"fmt"
)

type EventSender struct {

}

func (t *EventSender) Init(stub shim.ChaincodeStubInterface) peer.Response{
	err := stub.PutState("noevents",[]byte("0"))
	if err !=nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *EventSender) invoke(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	b, err := stub.GetState("noevents")
	if err != nil {
		return shim.Error(err.Error())
	}
	noevts, _ := strconv.Atoi(string(b))

	tosend := "Event"+ string(b)
	for _, s := range args {
		tosend = tosend + ","+ s
	}

	err = stub.PutState("noevents",[]byte(strconv.Itoa(noevts+1)))
	if err != nil{
		return shim.Error(err.Error())
	}

	// 利用此接口函数，来调用事件
	/*
	在fabric中，ChainCode除了可以主动查询账本，还可以在chainCode中发送事件。
	这个事件可以根据具体的业务逻辑来写
		- 比如，该事件可以是，当某个将以被Committer处被认证通过时，写入到区块的事件
	*/
	err = stub.SetEvent("eventsender",[]byte(tosend))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


func (t *EventSender) query(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	b ,err := stub.GetState("noevents")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	jsonResp := "{\"NoEvents\":\"" + string(b) + "\"}"
	return shim.Success([]byte(jsonResp))
}

func (t *EventSender) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")
}

func main(){
	err := shim.Start(new(EventSender))
	if err != nil{
		fmt.Printf("Error starting EventSendrt chaincode: %s",err)
	}
}