package gorm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli"
	"golang-example/cmd"
)

type (
	User struct {
		gorm.Model
		Name        string        `gorm:"TYPE:VARCHAR(255)"`
		Languages   []*Language   `gorm:"many2many:user_languages;"`
		CreditCards []*CreditCard `gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
		Projects    []*Project    `gorm:"many2many:user_projects;"`
	}

	Language struct {
		gorm.Model
		Name  string  `gorm:"TYPE:VARCHAR(255)"`
		Users []*User `gorm:"many2many:user_languages;"`
	}

	CreditCard struct {
		gorm.Model
		Number string `gorm:"TYPE:VARCHAR(255)"`
		UserID uint
	}
)

var devID = uuid.NewV4().String()

type (
	Project struct {
		gorm.Model
		UID     string   `gorm:"primary_key;type:blob;not null"`
		Name    string   `gorm:"TYPE:VARCHAR(255)"`
		Devices []Device `gorm:"foreignkey:ProjectUID;association_foreignkey:UID"`
		Users   []*User  `gorm:"many2many:user_projects;"`
	}

	Device struct {
		gorm.Model
		ProjectUID string
		UID        string `gorm:"primary_key;type:blob;not null"`
		Name       string `gorm:"TYPE:VARCHAR(255)"`
		Config     string `gorm:"type:blob"`
	}

	ModbusTCP struct {
		Name string
		Ip   string
	}
)

var (
	db  *gorm.DB
	err error
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gorm",
		Aliases: []string{"gorm"},

		Usage:    "Demonstration of gorm",
		Action:   gormAction,
		Category: "DB",
	})
}

func gormAction(c *cli.Context) error {
	InitDB()
	AutoMigrateUser()
	SetupUser()
	//Find()

	AutoMigrateProject()
	SetupProject()
	SetupDevices()
	FindProject()

	return nil
}

func InitDB() {
	sql.Register("sqlite3_with_extensions", &sqlite3.SQLiteDriver{
		Extensions: []string{
			"sqlite_userauth",
		},
	})

	db, err = gorm.Open("sqlite3",
		fmt.Sprintf("./misc/gorm/goexample.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha256", "admin", "123456"))

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
}

func AutoMigrateUser() {
	db.AutoMigrate(
		&User{},
		&Language{},
		&CreditCard{},
	)
}

func SetupUser() {
	var u User
	db.Where(&User{
		Name: "user1",
	}).First(&u)

	if db.HasTable(&User{}) && u.Name == "user1" {
		return
	}

	user1 := &User{
		Name: "user1",
	}

	user2 := &User{
		Name: "user2",
	}

	lang1 := &Language{
		Name: "Eng",
	}

	lang2 := &Language{
		Name: "Zh",
	}

	db.Create(lang1)
	db.Create(lang2)

	db.Create(user1).Association("Languages").Append([]*Language{lang1, lang2})
	db.Create(user2)

	var lan2 Language
	db.Where(&Language{
		Name: "Eng",
	}).First(&lan2).Association("Users").Append([]*User{user1, user2})

	card1 := &CreditCard{
		Number: "1",
	}

	card2 := &CreditCard{
		Number: "2",
	}

	var u2 User
	db.Where(&User{
		Name: "user1",
	}).First(&u2).Association("CreditCards").Append([]*CreditCard{card1, card2})
}

func Find() {
	// 查找指定 user 下的所有 Language
	var user User
	var langs []*Language
	tmpDB := db.Model(&user).Related(&langs, "Languages")

	var u User
	tmpDB.Preload("Languages").Where(&User{
		Name: "user1",
	}).Find(&u)

	// 查找指定 lang 下的所有 User
	var lan Language
	var users []*User
	tmpDB2 := db.Model(&lan).Related(&users, "Users")

	var l Language
	tmpDB2.Preload("Users").Where(&Language{
		Name: "Eng",
	}).Find(&l)

	// 查找指定 user 下的所有 CreditCard
	var user3 User
	//var creds []*CreditCard
	//tmpDB3 := db.Model(&user3).Related(&creds, "CreditCards")
	tmpDB3 := db.Model(&user3).Related(&user3.CreditCards)

	var u3 User
	tmpDB3.Preload("CreditCards").Where(&User{
		Name: "user1",
	}).Find(&u3)

	// 查找指定 user 下的所有 CreditCard
	var u4 User
	db.Find(&u4, "name = ?", "user1")
	db.Model(&u4).Association("CreditCards").Find(&u4.CreditCards)

	//var u5 User
	////db.First(&u4)
	//var c1 []*CreditCard
	//db.Model(&u5).Association("CreditCards").Find(&c1)
}

func AutoMigrateProject() {
	db.AutoMigrate(
		&Project{},
		&Device{},
	)
}

func SetupProject() {
	var p Project
	db.Where(&Project{
		Name: "p1",
	}).First(&p)
	if db.HasTable(&Project{}) && p.Name == "p1" {
		return
	}

	p1 := &Project{
		UID:  uuid.NewV4().String(),
		Name: "p1",
	}

	p2 := &Project{
		UID:  uuid.NewV4().String(),
		Name: "p2",
	}

	db.Create(p1)
	db.Create(p2)
}

func FindProject() {
	{
		//dev2 := Device{
		//	UID:  uuid.NewV4().String(),
		//	Name: "dev2",
		//}
		//
		//var p1 Project
		//db.Where(&Project{
		//	Name: "p1",
		//}).First(&p1).Association("Devices").Replace(Device{
		//	UID: devID,
		//}, dev2)

		var d Device
		db.Where(&Device{
			Name: "dev1",
		}).Find(&d)

		fmt.Println(d)
		d.Name = "dev2"
		db.Save(&d)
	}
	// 查找指定 Project 下的所有 Devices
	{
		var dev []Device
		tmpDB := db.Model(&Project{}).Related(&dev, "Devices")

		var p1 Project
		tmpDB.Preload("Devices").Where(&Project{
			Name: "p1",
		}).Find(&p1)

		//var m ModbusTCP
		//_ = json.Unmarshal([]byte(p1.Devices[0].Config), &m)

		fmt.Println(p1)
	}
	{
		var users []User
		tmpDB := db.Model(&Project{}).Related(&users, "Users")

		var p1 Project
		tmpDB.Preload("Users").Where(&Project{
			Name: "p1",
		}).Find(&p1).RecordNotFound()
		fmt.Println(p1)

		// 找到 user 后，再次查找 lang

	}
	{
		var users []User
		//users := []User{}
		tmpDB := getProjectPreload("Users", &users)

		var p1 Project
		tmpDB.Where(&Project{
			Name: "p1",
		}).Find(&p1)
		fmt.Println(p1)
	}
}

func SetupDevices() {
	var p Device
	db.Where(&Device{
		Name: "dev1",
	}).First(&p)
	if db.HasTable(&Project{}) && p.Name == "dev1" {
		return
	}

	modbus := &ModbusTCP{
		Name: "modbus",
		Ip:   "1.0.0.0",
	}

	data, _ := json.Marshal(modbus)

	dev1 := Device{
		UID:    devID,
		Name:   "dev1",
		Config: string(data),
		//Config: &ModbusTCP{
		//	Name: "modbus",
		//	Ip:   "1.0.0.0",
		//},
	}

	var p1 Project
	db.Where(&Project{
		Name: "p1",
	}).First(&p1).Association("Devices").Append([]Device{dev1})

	var u1 User
	db.Where(&User{
		Name: "user1",
	}).First(&u1).Association("Projects").Append([]Project{p1})
}

func getProjectPreload(assoc string, subType interface{}) *gorm.DB {
	if len(assoc) == 0 {
		return nil
	}

	//var p Project
	//var tmpNets []model.Network
	//tmpDB := s.db.Model(&model.Project{}).Related(&tmpNets, config.DBAssocNetworks)
	tmpDB := db.Model(&Project{}).Related(subType, assoc)

	return tmpDB.Preload(assoc)
}
