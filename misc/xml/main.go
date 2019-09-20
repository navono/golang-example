package xml

import (
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/urfave/cli"
	"golang-example/cmd"
	"strconv"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "xml",
		Aliases: []string{"x"},

		Usage:    "Demonstration of xml operations",
		Category: "data",
		Subcommands: []cli.Command{
			{
				Name:   "basic",
				Usage:  "basic example",
				Action: basicXmlAction,
			},
			{
				Name:   "file",
				Usage:  "file example",
				Action: fileXmlAction,
			},
		},
	})
}

var data = []byte(`
<a>
  <b>1</b>
</a>`)

// other example see: https://github.com/clbanning/mxj/blob/master/examples
func basicXmlAction(c *cli.Context) error {
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("new  err:", err)
		return err
	}

	b, err := m.ValueForPath("a.b")
	if err != nil {
		fmt.Println("value err:", err)
		return err
	}

	b, err = appendElement(b, 2)
	if err != nil {
		fmt.Println("append err:", err)
		return err
	}

	// Create the new value for 'b' as a map
	// and update 'm'.
	// We should probably have an UpdateValueForPath
	// method just as there is ValueForPath/ValuesForPath
	// methods.
	val := map[string]interface{}{"b": b}
	n, err := m.UpdateValuesForPath(val, "a.b")
	if err != nil {
		fmt.Println("update err:", err)
		return err
	}
	if n == 0 {
		fmt.Println("err: a.b not updated, n =", n)
		return err
	}

	x, err := m.XmlIndent("", "  ")
	if err != nil {
		fmt.Println("XmlIndent err:", err)
		return err
	}
	fmt.Println(string(x))
	return nil
}

func fileXmlAction(c *cli.Context) error {
	maps, err := mxj.NewMapsFromXmlFile("./misc/xml/Project.xml")
	if err != nil {
		return err
	}
	project := maps[0]

	fmt.Println(project)
	return nil
}

func appendElement(v interface{}, n int) (interface{}, error) {
	switch v.(type) {
	case string:
		v = []interface{}{v.(string), strconv.Itoa(n)}
	case []interface{}:
		v = append(v.([]interface{}), interface{}(strconv.Itoa(n)))
	default:
		// capture map[string]interface{} value, simple element, etc.
		return v, fmt.Errorf("invalid type")
	}
	return v, nil
}
