package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type EduCertify struct {

}

var BackGroundNo int = 0
var RecordNO int = 0

type School struct {
	Name           string
	Location       string
	Address        string
	PriKey         string
	PubKey         string
	StudentAddress []string
}

type Student struct {
	Name         string
	Address      string
	BackgroundId []int
}

// 学历信息，一般在学生在一个学校完成信息后完成，所以，一个学生可能对应多个BackGround.是一种一对多的关系。
type Background struct {
	Id int
	ExitTime int64
	Status string  // 0:毕业 1：退学
}


//   修改记录，升学、退学、入学，都会对学生的BackGround 做一个修改，而每一个修改对应一个Record。
//   这里的Record，类似于数据库中的undo\redo 日志
type Record struct {
	Id              int
	SchoolAddress   string
	StudentAddress  string
	SchoolSign      string
	ModifyTime      int64
	ModifyOperation string // 0:正常毕业 1：退学 2:入学
}

/*
 * 区块链网络实例化“diploma”智能合约时会调用该方法
 */
func (e *EduCertify) Init(stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Success(nil)
}

/*
 *客户端发起执行“diploma”智能合约时会调用Invoke方法
 *
*/
func (e *EduCertify) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	// 获取请求调用智能合约的方法和参数
	function, args := stub.GetFunctionAndParameters()
	// 根据不同的方法，路由至不同的处理函数
	if function == "createSchool"{
		return e.createSchool(stub, args)
	}else if function == "createStudent"{
		return e.createStudent(stub, args)
	}
}

/*
 * 在区块链网络中，来创建一个School
 * 传入参数： args[0] 学校名称 args[1] 学校所在位置
*/
func (e *EduCertify) createSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response  {
	// 判断参数是否合法
	if len(args) != 2{
		return shim.Error("Incorrect number of arguments. Excepting 2")
	}

	//根据对应的参数，来创建school
	var school School
	var schoolBytes []byte
	var stuAddress []string
	var address, priKey, pubKey = GetAddress()

	school = School{Name:args[0],Location:args[1],Address:address,PriKey:priKey,PubKey:pubKey,StudentAddress:stuAddress}
	err := writeSchool(stub, school)
	if err != nil{
		shim.Error("Error write school")
	}

	schoolBytes, err = json.Marshal(&school)
	if err != nil{
		return shim.Error("Error retrieving schlloBytes")
	}
	return shim.Success(schoolBytes)
}

/*writeSchool()
 *将School实体存入到区块链中
 *传入 school 实例
*/

func writeSchool(stub shim.ChaincodeStubInterface, school School) error{
	schoolBytes, err := json.Marshal(school)
	if err != nil{
		return errors.New("Error Marshal school")
	}
	err = stub.PutState(school.Address,schoolBytes)
	if err != nil{
		return errors.New("PutState Error" + err.Error())
	}
	return nil
}




/*getAddress(),获取地址
 *
 *返回账户地址、私钥、公钥
 *
*/
func GetAddress() (string, string, string){
	var address, priKey, pubKey string
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err !=nil{
		return "", "", ""
	}

	//使用md5以及base64编码，生成对应的账户地址和公私钥
	h := md5.New()
	//因为生成的地址需要用在url中，需要使用URLEncoding，进行base64编码
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))

	address = hex.EncodeToString(h.Sum(nil))
	priKey = address + "1"
	pubKey = address + "2"
	return address, priKey, pubKey
}

/*
 *createStudent()函数，
 *args[0] 学生姓名
*/
func (e *EduCertify) createStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	// 判断参数是否合法
	if len(args) != 1{
		return shim.Error("Incorrect Number of arguments.Excepting 1")
	}

	// 为创建Student定义自变量
	var student Student
	var stuBytes []byte
	var stuAddress string
	var backgdId []int


	stuAddress, _, _ = GetAddress()
	student = Student{Name:args[0],Address:stuAddress,BackgroundId:backgdId}
	err := writeStudent(stub,student)
	if err != nil{
		return shim.Error("Failed to write into BlockChain")
	}
	stuBytes,err = json.Marshal(&student)
	if nil != nil{
		return shim.Error("Error retriving student")
	}
	return shim.Success(stuBytes)
}

/*
 *writeStudent
 *
*/
func writeStudent(stub shim.ChaincodeStubInterface, student Student) error {
	studentBytes, err := json.Marshal(student)
	if err != nil {
		return err
	}
	err = stub.PutState(student.Address,studentBytes)
	if err != nil{
		return errors.New("Failed to write into BlcokChain")
	}
	return nil
}