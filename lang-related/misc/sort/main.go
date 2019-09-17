package sort

import (
	"fmt"
	"github.com/urfave/cli"
	"golang-example/cmd"
	"sort"
	"strings"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "sort",
		Aliases: []string{"s"},

		Usage:    "Demonstration of slice sort operations",
		Action:   sortAction,
		Category: "lang-misc",
	})
}

type people struct {
	Id  string
	Age uint16
}

type peopleSlice []people

func (p peopleSlice) Len() int {
	return len(p)
}

func (p peopleSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p peopleSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func sortAction(c *cli.Context) error {
	students := []people{
		{
			Id:  "1",
			Age: 10,
		}, {
			Id:  "2",
			Age: 5,
		}, {
			Id:  "3",
			Age: 8,
		},
	}

	sort.Sort(peopleSlice(students))
	fmt.Println(students)

	idx := sort.Search(len(students), func(i int) bool {
		return strings.Compare(students[i].Id, "3") == 0
	})

	fmt.Println(idx)

	return nil
}
