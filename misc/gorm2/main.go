package gorm2

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gorm2",
		Aliases: []string{"gorm2"},

		Usage:    "Demonstration of gorm",
		Action:   gormAction,
		Category: "DB",
	})
}

func gormAction(c *cli.Context) error {
	initDB()
	createProject()
	associateNet()
	associateDevice()

	query()

	defer db.Close()
	return nil
}

var (
	db  *gorm.DB
	err error
)

type (
	Project struct {
		gorm.Model
		Name     string
		Networks []Network
		Devices  []Device
	}

	Network struct {
		gorm.Model
		ProjectID uint
		Name      string
		Devices   []Device `gorm:"many2many:net_dev;"`
	}

	Device struct {
		gorm.Model
		ProjectID uint
		Name      string
		Networks  []Network `gorm:"many2many:net_dev;"`
	}
)

func initDB() {
	sql.Register("sqlite3_with_extensions", &sqlite3.SQLiteDriver{
		Extensions: []string{
			"sqlite_userauth",
		},
	})

	db, err = gorm.Open("sqlite3",
		fmt.Sprintf("./misc/gorm2/db2.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha256", "admin", "123456"))

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	db.AutoMigrate(
		&Project{},
		&Network{},
		&Device{},
	)
}

func createProject() {
	var u Project
	db.Where(&Project{
		Name: "p1",
	}).First(&u)

	if db.HasTable(&Project{}) && u.Name == "p1" {
		return
	}

	p1 := &Project{
		Name: "p1",
	}

	db.Create(&p1)
}

func associateNet() {
	net := Network{
		Name: "net1",
	}

	var n Network
	db.Where(&net).First(&n)
	if db.HasTable(&Network{}) && n.Name == "net1" {
		return
	}

	var p Project
	db.Where(&Project{
		Name: "p1",
	}).Find(&p).Association("Networks").Append([]Network{net})
}

func associateDevice() {
	dev := Device{
		Name: "dev1",
	}

	var d Device
	db.Where(&dev).First(&d)
	if db.HasTable(&Network{}) && d.Name == "dev1" {
		return
	}

	// Project 关联 Devices
	var p Project
	db.Where(&Project{
		Name: "p1",
	}).Find(&p).Association("Devices").Append([]Device{dev})

	// 找到 Device 需要关联的 Network
	var n Network
	db.Where(&Network{
		Name: "net1",
	}).Find(&n)

	// 再对已存在的 Device 关联 Network
	var d1 Device
	db.Where(&Device{
		Name: dev.Name,
	}).Find(&d1).Association("Networks").Append([]Network{n})
}

func query() {
	// 查找 project 的所有 devices
	p := Project{
		Name: "p1",
	}
	var tmpDev []Device
	db.Model(&Project{}).Related(&tmpDev, "Devices").Preload("Devices").Where(&p).Find(&p)
	fmt.Println(p.Devices)

	// 查找 project 的所有 network
	var tmpNet []Network
	db.Model(&Project{}).Related(&tmpNet, "Networks").Preload("Networks").Where(&Project{
		Name: "p1",
	}).Find(&p)
	fmt.Println(p.Networks)

	// 查找 device 的 networks
	d := Device{
		Name: "dev1",
	}
	db.Model(&Device{}).Related(&tmpNet, "Networks").Preload("Networks").Where(&d).Find(&d)
	fmt.Println(d)

	// 查找 network 的 devices
	n := Network{
		Name: "net1",
	}
	db.Model(&Network{}).Related(&tmpDev, "Devices").Preload("Devices").Where(&n).Find(&n)
	fmt.Println(n)
}
