package dialog

import (
	"fmt"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "dialog",
		Aliases: []string{"dl"},

		Usage:    "Demonstration of gorm",
		Action:   dialogAction,
		Category: "DB",
	})
}

func dialogAction(c *cli.Context) error {
	// _ = dialog.Message("%s", "Do you want to continue?").Title("Are you sure?").YesNo()

	/* Note that spawning a dialog from a non-graphical app like this doesn't
	** quite work properly in OSX. The dialog appears fine, and mouse
	** interaction works but keypresses go straight through the dialog.
	** I'm guessing it has something to do with not having a main loop? */
	dialog.Message("%s", "Please select a file").Title("Hello world!").Info()
	// file, err := dialog.File().Title("Save As").Filter("All Files", "*").Save()
	file, err := dialog.File().Title("Save As").Filter("All Files", "*").Load()
	fmt.Println(file)
	fmt.Println("Error:", err)
	dialog.Message("You chose file: %s", file).Title("Goodbye world!").Error()
	d, err := dialog.Directory().Title("Now find a dir").Browse()

	nl := strings.Split(d, "\\")

	name := nl[len(nl)-1]
	fmt.Println(name)

	return nil
}
