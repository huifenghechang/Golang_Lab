package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"strconv"
	"encoding/json"
	"time"
	"fmt"
)

type SimpleChainCode struct {

}

type CenterBank struct {
	Name string
	TotalNumber int
	RestNumber int
	ID int
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
	ToType string
	ToID int
	Time int64
	Number int
	ID int
}

var bankNO int = 0
var cpNO int = 0
var transactionNO int = 0

/*
	init 接口，用来初始化相关参数
	 request 参数 ：args[0] 银行名称； args[1]  初始化发布金额
	 response参数 ：{"Name":"XXX","TotalNumber":"0","RestNumber":"0","ID":"XX"}
*/

func (s *SimpleChainCode) Init(stub *shim.ChaincodeStub,function string, args []string) ([]byte, error){

	if len(args) != 2{
		shim.Error("Incorrect number of arguments.Excepting 2")
		return []byte("error args"), errors.New("Error Arguments")
	}

	var centerBkName string
	var initalNumber int
	centerBkName = args[0]
	initalNumber, err := strconv.Atoi(args[1])
	if err != nil{
		return []byte("Error arguments"), errors.New("Error arguments.Excepting an interger")
	}

	// 创建一个CenterBank实例，并将其存储到区块链中
	centerBk := CenterBank{Name:centerBkName,TotalNumber:initalNumber,RestNumber:0,ID:0}
	err = writeBank(stub,centerBk)

	if err != nil{
		return nil,errors.New("Failed to store in blockChain"+err.Error())
	}
	return nil,nil
}


// Invoke() 函数

func (t *SimpleChainCode) Invoke (stub *shim.ChaincodeStub, function string, args []string) ([]byte, error)  {

	if function == "createBank"{
		t.createBank(stub, args)
	}else if function == "createCompany"{
		return t.createCompany(stub,args)
	}else if function == "issueCoin"{
		return t.issueCoin(stub,args)
	}else if function == "issueCoinToBank"{
		return t.issueCoinToBank(stub,args)
	}

}


// request 参数： args[0] 银行名称
func (t *SimpleChainCode) createBank(stub *shim.ChaincodeStub, args []string) (string) {
	if len(args) != 1 {
		return "Incorrect Number of arguments. Excepting 1"
	}
	var bankName string
	var bank Bank
	bankName = args[0]
	bank = Bank{Name:bankName,TotalNumber:0,RestNumber:0,ID:bankNO}
	bankAsBytes, err := json.Marshal(bank)
	if err != nil{
		return "Failed in marshal bank to json"
	}

	err = stub.PutState(bankName,bankAsBytes)
	if err != nil {
		return "Failed to store bank into BlockChain"
	}

	bankNO = bankNO + 1
	return "create Bank Success!"
}


// request 参数 args[0] 公司名称
func (t *SimpleChainCode) createCompany(stub *shim.ChaincodeStub, args []string) ([]byte, error){
	if len(args) != 1{
		return nil, errors.New("Incorrect number of arguments. Excepting 1")
	}

	var company Company
	company = Company{Name:args[0],Number:0,ID:cpNO}
	companyAsBytes,err := json.Marshal(company)
	if err != nil{
		return nil, errors.New("Error retrieving cpBytes")
	}

	err = writeCompany(stub, company)
	if err != nil {
		return nil, errors.New("Write Error" + err.Error())
	}
	cpNO = cpNO + 1
	return companyAsBytes,nil
}


// issueCoin 发行货币，args[0] 再次发行货币数额
func (t *SimpleChainCode) issueCoin(stub *shim.ChaincodeStub, args []string) ([]byte, error){
	if len(args) != 1{
		return nil, errors.New("Incorrect number of arguments. Excepting 1")
	}

	var issueNum int
	var centerBank CenterBank
	var tsBytes []byte

	issueNum, err := strconv.Atoi(args[0])
	if err != nil{
		return nil, errors.New("Excepting an Interger")
	}
	centerBank,_, err = getCenterBank(stub)
	if err != nil{
		return nil, errors.New("get Errors"+err.Error())
	}

	//更改央行信息，进行发行积分
	centerBank.TotalNumber = centerBank.TotalNumber + issueNum
	centerBank.RestNumber = centerBank.TotalNumber + issueNum

	err = writeCenterBank(stub,centerBank)
	if err != nil{
		return nil, errors.New("write Error !" + err.Error())
	}

	transaction := Transaction{FromType:0,FormID:0,ToID:0,Time:time.Now().Unix(),Number:issueNum,ID:transactionNO}
	err = writeTransaction(stub,transaction)
	if err != nil {
		return nil, errors.New("write Error" + err.Error())
	}

	tsBytes, err = json.Marshal(transaction)
	if err != nil{
		return nil, errors.New("Error marshal Transaction")
	}

	transactionNO = transactionNO + 1
	return tsBytes, nil
}


// args[0] 商业银行ID
// args[1] 转账数额

/*
	编程基本思路：
	首先，从区块链中，读取商业银行，读取中央银行的值；
	其次，将中央银行的值减去、商业银行的值增加
	创建一笔交易，将交易存入区块链中。


	编程之问：
	- 在这个函数中，centerBank，不需要实例化，可直接将从区块链中读取的数据赋值给该变量。
*/

func (t *SimpleChainCode) issueCoinToBank(stub *shim.ChaincodeStub, args []string)([]byte, error){
	if len(args) != 2{
		return nil, errors.New("Incorrect Number of arguments. Excepting 2")
	}

	var centerBank CenterBank
	var bank Bank
	var bankID int
	var issuseNum int

	bankID, err := strconv.Atoi(args[0])
	if err != nil{
		return nil, errors.New("Arguments Incorrect.Excepting an Interger")
	}

	issuseNum, err = strconv.Atoi(args[1])
	if err != nil{
		return nil,errors.New("Arguments Incorrect.Excepting an interger")
	}

	bank, _ , err = getBankByID(stub,bankID)
	if err != nil{
		return nil, errors.New("Error getBank")
	}

	centerBank, _, err = getCenterBank(stub)

	if err != nil{
		return nil, errors.New("Error get CenterBank")
	}

	// 进行发行数字积分操作
	if issuseNum > centerBank.RestNumber{
		return nil,errors.New("the RestNumber is litter than issueseNum")
	}

	bank.RestNumber = bank.RestNumber + issuseNum
	bank.TotalNumber = bank.TotalNumber + issuseNum
	centerBank.RestNumber = centerBank.RestNumber - issuseNum

	//若出现错误，则将操作回滚
	err = writeCenterBank(stub,centerBank)
	if err != nil{
		bank.RestNumber = bank.RestNumber - issuseNum
		bank.TotalNumber = bank.TotalNumber - issuseNum
		centerBank.RestNumber = centerBank.RestNumber + issuseNum
		return  nil, errors.New("write errors"+ err.Error())
	}

	err = writeBank(stub,bank)
	if err != nil{
		bank.RestNumber = bank.RestNumber - issuseNum
		bank.TotalNumber = bank.TotalNumber - issuseNum
		centerBank.RestNumber = centerBank.RestNumber + issuseNum
		err = writeCenterBank(stub,centerBank)
		if err != nil{
			return nil, errors.New("Roll Down errors."+err.Error())
		}
		return nil, err
	}

	// 若发行成功，则创建一个交易实体，记录该笔交易
	transaction := Transaction{FromType:0,FormID:0,ToType:1,ToID:bankID,Time:time.Now().Unix(),Number:issuseNum,ID:transactionNO}
	err = writeTransaction(stub,transaction)
	if err != nil{
		return nil, errors.New("Transaction Write Error"+ err.Error())
	}

	tsBytes, err := json.Marshal(transaction)
	if err != nil{
		return nil, errors.New("Error Unmarshal Transaction")
	}
	transactionNO = transactionNO + 1
	return tsBytes, nil
}


/*
写一系列的get 和 set 函数， 用于从区块链中存入和读取数据。
*/
func getCenterBank(stub *shim.ChaincodeStub) (CenterBank, []byte, error){
	var centerBank CenterBank
	cbBytes, err := stub.GetState("centerBank")
	if err != nil{
		fmt.Printf("Error retriving cbBytes")
	}
	return centerBank, cbBytes,nil
}

/*
根据bankID，来获取对应的值
*/

func getBankByID(stub *shim.ChaincodeStub, bankID int)(Bank, []byte, error){

}

func writeCenterBank(stub *shim.ChaincodeStub,centerBank CenterBank) (error) {
	cbBytes, err := json.Marshal(&centerBank)
	if err != nil {
		return err
	}
	err = stub.PutState("centerBank", cbBytes)
	if err != nil {
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}

func writeBank(stub *shim.ChaincodeStub,bank Bank)(error)  {
	bankAsbyte, err := json.Marshal(&bank)
	if err != nil{
		return errors.New("Failed to tranfer struct type to json")
	}
	err = stub.PutState(bank.Name+strconv.Itoa(bank.ID),[]byte(bankAsbyte))
	return err
}


func writeCompany(stub *shim.ChaincodeStub, company Company) ( error){
	companyAsBytes,err := json.Marshal(company)
	if err != nil{
		return  errors.New("Error retrieving cpBytes")
	}

	err = stub.PutState(company.Name, companyAsBytes)
	if err != nil {
		return errors.New("Write Error" + err.Error())
	}
	return nil
}

func writeTransaction(stub *shim.ChaincodeStub, transaction Transaction) (error){
	tsBytes, err := json.Marshal(transaction)
	if err != nil{
		return errors.New("Error marshal transaction")
	}

	tsId := strconv.Itoa(transaction.ID)

	err = stub.PutState("transaction"+tsId,tsBytes)
	if err != nil {
		return errors.New("PutState Error!" + err.Error())
	}
	return nil
}

func main(){
}
