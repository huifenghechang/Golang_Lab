package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	/*fmt.Println("I am HandleFileList -- -- 001 ")
	if strings.Index(
		request.URL.Path,prefix) != 0{
		return errors.New("path must start"+
			" with " + prefix)
	}

	path := request.URL.Path[len(prefix):]*/
	fmt.Println("I am HandleFileList -- -- 001 ")
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
