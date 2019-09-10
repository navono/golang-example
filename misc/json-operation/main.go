package json_operation

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/urfave/cli"
	"golang-example/cmd"
	"io/ioutil"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "json",
		Aliases: []string{"j"},

		Usage:    "Demonstration of json operations",
		Action:   jsonAction,
		Category: "data",
	})
}

func jsonAction(c *cli.Context) error {
	fileContent, _ := readFile("./misc/json-operation/test.json")
	jsonParsed, err := gabs.ParseJSON(fileContent)

	if err != nil {
		panic(err)
	}

	_ = jsonParsed.Delete("nets", "1")

	//nets := jsonParsed.S("nets")
	//_ = nets.Delete("1")

	fmt.Println(jsonParsed)
	return nil
}

func readFile(filename string) ([]byte, error) {
	if len(filename) == 0 {
		return nil, errors.New("filename empty")
	}

	return ioutil.ReadFile(filename)
}
