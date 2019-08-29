package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
	"golang-example/cmd"
)

type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

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
	InitSqlite()
	AutoMigrate()
	Setup()
	Find()

	return nil
}

func InitSqlite() {
	db, err = gorm.Open("sqlite3",
		fmt.Sprintf("config.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha256", "admin", "123456"))

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
}

func AutoMigrate() {
	db.AutoMigrate(
		&User{},
		&Language{},
	)
}

func Setup() {
	if db.HasTable(&User{}) {
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
}

func Find() {
	var user User
	var langs []*Language
	tmpDB := db.Model(&user).Related(&langs, "Languages")

	var u User
	tmpDB.Preload("Languages").Where(&User{
		Name: "user2",
	}).Find(&u)

	var lan Language
	var users []*User
	tmpDB2 := db.Model(&lan).Related(&users, "Users")

	var l Language
	tmpDB2.Preload("Users").Where(&Language{
		Name: "Eng",
	}).Find(&l)
}
