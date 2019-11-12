package gorm3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"

	"github.com/satori/go.uuid"

	"golang-example/cmd"
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gorm3",
		Aliases: []string{"gorm3"},

		Usage:    "Demonstration of gorm",
		Action:   gorm3Action,
		Category: "DB",
	})
}

var (
	db *gorm.DB
)

func gorm3Action(c *cli.Context) error {
	initDB()
	fillData()

	var nt NetworkTemplate
	db.Model(&NetworkTemplate{}).First(&nt)
	fmt.Println(nt)

	return nil
}

func initDB() {
	db, _ = gorm.Open("sqlite3",
		fmt.Sprintf("./misc/gorm3/db.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha256", "admin", "123456"))

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	db.AutoMigrate(&NetworkTemplate{})
}

func fillData() {
	var n NetworkTemplate
	db.Where(NetworkTemplate{
		Name: "Net3",
	}).First(&n)
	if db.HasTable(&NetworkTemplate{}) && n.Name == "Net3" {
		return
	}

	net3 := net3{
		Addrs: []netElement{
			{
				Name:          "控制网A",
				Ip:            "172.20.*.*",
				CanBeDisabled: true,
			}, {
				Name:          "控制网B",
				Ip:            "172.21.*.*",
				CanBeDisabled: true,
			}, {
				Name:          "信息网A",
				Ip:            "172.30.*.*",
				CanBeDisabled: true,
			}, {
				Name:          "信息网B",
				Ip:            "172.31.*.*",
				CanBeDisabled: true,
			},
		},
		NTPServer: []netElement{
			{
				Name:          "时间服务器",
				Ip:            "*.*.*.254",
				CanBeDisabled: true,
			},
		},
	}

	config, _ := json.Marshal(&net3)
	net := NetworkTemplate{
		UID:    uuid.NewV4().String(),
		Name:   "Net3",
		Type:   1,
		Config: string(config),
	}

	db.Create(&net)
}
