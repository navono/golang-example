package filesystem

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

var appFs *afero.Afero

func aferoAction(c *cli.Context) error {
	appFs = &afero.Afero{Fs: afero.NewMemMapFs()}

	f, _ := appFs.Create("/a")
	_, _ = f.Write([]byte("test content"))

	appFs.Walk("/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)

		return nil
	})

	return nil
}
