package filehandling

import (
	"fmt"
	"io/ioutil"
	"os"
)

func testReadRelativeFile() {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "\\IO\\test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func init() {
	fmt.Println()
	fmt.Println("===> enter IO package")

	testReadRelativeFile()

	fmt.Println("<=== exit IO package")
	fmt.Println()
}
