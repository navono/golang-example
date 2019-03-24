package filehandling

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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

func testBufferedRead() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "\\IO\\test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

func testReadFileLineByLine() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "\\IO\\test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	fmt.Println()
	fmt.Println("===> enter IO package")

	// testReadRelativeFile()
	// testBufferedRead()
	testReadFileLineByLine()

	fmt.Println("<=== exit IO package")
	fmt.Println()
}
