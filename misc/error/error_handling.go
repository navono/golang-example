package errorhandling

import (
	"fmt"
	"os"
	"path/filepath"
)

func openFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.Name(), "opened successfully")
}

func openFileWithErrorAssert(filename string) {
	f, err := os.Open(filename)
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}

	fmt.Println(f.Name(), "opened successfully")
}

// var ErrBadPattern = errors.New("syntax error in pattern")
func getFileList() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter error_handling package")

	// openFile("./test.txt")
	// openFileWithErrorAssert("./test.txt")

	// getFileList()

	// testCustomError()
	testTypedError()

	fmt.Println("<=== exit error_handling package")
	fmt.Println()
}
