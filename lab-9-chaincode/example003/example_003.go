package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"errors"
	"strconv"
	"encoding/json"
)

// 设置全局变量
var bankNO int = 0
var cpNO int = 0
var transactionNO = 0

type SimpleChaincode struct {

}

type CenterBank struct {
	Name string
	TotalNumber int
	RestNumber int
}

type Bank struct {
	Name string
	TotalNumber int
	RestNumber int
	ID int
}

type Company struct {
	Name string
	Number int
	ID int
}

type Transaction struct {
	FromType int
	FormID int
	ToType int
	ToID int
	Time int64
	Number int
	ID int
}

func main(){
	err := shim.Start(new(SimpleChaincode))
	if err != nil{
		fmt.Printf("Error starting Simple ChainCode",err)
	}
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	if len(args) != 2{
		return nil, errors.New("Incorrect number of arguments. Excepting 2")
	}

	// 在写函数的时候，一定要遵循先定义变量，后写业务代码的习惯。
	// 这样有两个好处，一是代码思路很明确， 二是代码的可读性高。
	var totalNumber int
	var centerBank CenterBank
	var cbBytes []byte
	totalNumber, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Excepting interger value for asset holding")
	}
	centerBank = CenterBank{Name:args[0], TotalNumber:totalNumber, RestNumber:0}

	err = writeCenterBank(stub, centerBank)
	if err != nil {
		return nil, errors.New("Write Error" + err.Error())
	}

	cbBytes, err = json.Marshal(&centerBank)
	if err != nil {
		return nil,errors.New("Marshal Error" + err.Error())
	}
	return cbBytes, nil
}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub,function string,args []string) ([]byte, error)  {
	if function == "createBank" {
		t.createBank(stub,args)
	}

}

func (t *SimpleChaincode) createBank(stub *shim.ChaincodeStub, args []string) ([]byte,error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Excepting 1")
	}
	var bank Bank
	var bankBytes []byte

	bank = Bank{Name:args[0],TotalNumber:0, RestNumber:0, ID:bankNO}
	err := writeBank(stub,bank)
	if err != nil {
		return nil,errors.New("Write Error" + err.Error())
	}
	bankBytes, err = json.Marshal(bank)
	if err != nil {
		return nil,errors.New("Error retrieving cbBytes")
	}
	return bankBytes,nil

}

// 对于经常性的操作，封装成函数。
func writeCenterBank(stub *shim.ChaincodeStub, centerBank CenterBank ) (error) {
	cbBytes, err := json.Marshal(centerBank)
	if err != nil {
		return err
	}
	err = stub.PutState("centerBank", cbBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}

// 对于经常性的操作，封装成函数。
func writeBank(stub *shim.ChaincodeStub, bank Bank ) (error) {
	cbBytes, err := json.Marshal(bank)
	if err != nil {
		return err
	}
	err = stub.PutState(bank.Name, cbBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}