package gabs

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
	fileContent, _ := readFile("./misc/gabs/test.json")
	jsonParsed, err := gabs.ParseJSON(fileContent)

	if err != nil {
		panic(err)
	}

	//_ = jsonParsed.Delete("nets", "1")

	//nets := jsonParsed.S("nets")
	//_ = nets.Delete("1")

	modifyJson, _ := gabs.ParseJSON([]byte(`
{
  "projectID": "175e9551-a268-4f36-9ece-57dfe273a2dd",
  "nets": [
    {
      "name": "test_net_modify",
      "id": "5e45f95f-b039-4b10-b737-7663bf2581c3",
      "type": 1,
      "time_sources": [
        "1",
        "2"
      ]
    }
  ]
}
`))

	_ = jsonParsed.MergeFn(modifyJson, func(destination, source interface{}) interface{} {
		// 从 destination 找到 source 相应中的数据，然后替换到 destination 中
		switch u := destination.(type) {
		case string:
			fmt.Println(u)
		case *[]interface{}:
			fmt.Println(u)
		}

		return source
	})

	fmt.Println(jsonParsed)
	return nil
}

func readFile(filename string) ([]byte, error) {
	if len(filename) == 0 {
		return nil, errors.New("filename empty")
	}

	return ioutil.ReadFile(filename)
}
