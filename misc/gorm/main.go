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
	Name        string        `gorm:"TYPE:VARCHAR(255)"`
	Languages   []*Language   `gorm:"many2many:user_languages;"`
	CreditCards []*CreditCard `gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
}

type Language struct {
	gorm.Model
	Name  string  `gorm:"TYPE:VARCHAR(255)"`
	Users []*User `gorm:"many2many:user_languages;"`
}

type CreditCard struct {
	gorm.Model
	Number string `gorm:"TYPE:VARCHAR(255)"`
	UserID uint
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
	//sql.Register("sqlite3_with_extensions", &sqlite3.SQLiteDriver{
	//	Extensions: []string{
	//		"sqlite_userauth",
	//	},
	//})

	db, err = gorm.Open("sqlite3",
		fmt.Sprintf("goexample.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=sha256", "admin", "123456"))

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
}

func AutoMigrate() {
	db.AutoMigrate(
		&User{},
		&Language{},
		&CreditCard{},
	)
}

func Setup() {
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
