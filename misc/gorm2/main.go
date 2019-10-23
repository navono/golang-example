package gorm2

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"

	"golang-example/cmd"
	"golang-example/misc/gorm2/model"
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
		&model.Project{},
		&model.Network{},
		&model.Device{},
	)
}

func createProject() {
	var u model.Project
	db.Where(&model.Project{
		Name: "p1",
	}).First(&u)

	if db.HasTable(&model.Project{}) && u.Name == "p1" {
		return
	}

	p1 := &model.Project{
		Name: "p1",
	}

	db.Create(&p1)
}

func associateNet() {
	net := model.Network{
		Name: "net1",
	}

	var n model.Network
	db.Where(&net).First(&n)
	if db.HasTable(&model.Network{}) && n.Name == "net1" {
		return
	}

	var p model.Project
	db.Where(&model.Project{
		Name: "p1",
	}).Find(&p).Association("Networks").Append([]model.Network{net})
}

func associateDevice() {
	dev := model.Device{
		Name: "dev1",
	}

	var d model.Device
	db.Where(&dev).First(&d)
	if db.HasTable(&model.Network{}) && d.Name == "dev1" {
		return
	}

	// Project 关联 Devices
	var p model.Project
	db.Where(&model.Project{
		Name: "p1",
	}).Find(&p).Association("Devices").Append([]model.Device{dev})

	// 找到 Device 需要关联的 Network
	var n model.Network
	db.Where(&model.Network{
		Name: "net1",
	}).Find(&n)

	// 再对已存在的 Device 关联 Network
	var d1 model.Device
	db.Where(&model.Device{
		Name: dev.Name,
	}).Find(&d1).Association("Networks").Append([]model.Network{n})
}

func query() {
	// 查找 project 的所有 devices
	p := model.Project{
		Name: "p1",
	}
	var tmpDev []model.Device
	db.Model(&model.Project{}).Related(&tmpDev, "Devices").Preload("Devices").Where(&p).Find(&p)
	fmt.Println(p.Devices)

	// 查找 project 的所有 network
	var tmpNet []model.Network
	db.Model(&model.Project{}).Related(&tmpNet, "Networks").Preload("Networks").Where(&model.Project{
		Name: "p1",
	}).Find(&p)
	fmt.Println(p.Networks)

	// 查找 device 的 networks
	d := model.Device{
		Name: "dev1",
	}
	db.Model(&model.Device{}).Related(&tmpNet, "Networks").Preload("Networks").Where(&d).Find(&d)
	fmt.Println(d)

	// 查找 network 的 devices
	n := model.Network{
		Name: "net1",
	}
	db.Model(&model.Network{}).Related(&tmpDev, "Devices").Preload("Devices").Where(&n).Find(&n)
	fmt.Println(n)
}
