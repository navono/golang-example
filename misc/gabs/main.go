package gabs

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/urfave/cli"
	"golang-example/cmd"
	"io/ioutil"
)

type net struct {
	Name        string        `json:"name,omitempty"`
	Id          string        `json:"id,omitempty"`
	Type        int32         `json:"type,omitempty"`
	Segments    []*netSegment `json:"segments,omitempty"`
	TimeSources []string      `json:"time_sources,omitempty"`
}

type netSegment struct {
	Name     string `json:"name,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gabs",
		Aliases: []string{"j"},

		Usage:    "Demonstration of gabs operations",
		Action:   jsonAction,
		Category: "data",
	})
}

func jsonAction(c *cli.Context) error {
	fileContent, _ := readFile("./misc/gabs/test.gabs")
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

	var n []net
	_ = mapstructure.Decode(modifyJson.S("nets").Data(), n)

	fmt.Println(jsonParsed)
	return nil
}

func readFile(filename string) ([]byte, error) {
	if len(filename) == 0 {
		return nil, errors.New("filename empty")
	}

	return ioutil.ReadFile(filename)
}
