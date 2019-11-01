package gorm_multiDB

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

var (
	db1Name = "db1.db"
	db2Name = "db2.db"
	db1     *gorm.DB
	db2     *gorm.DB
	id      = "test"
)

type (
	A struct {
		ID   string
		Name string
	}
	B struct {
		ID   string
		Type string
	}
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gorm-m",
		Aliases: []string{"gm"},

		Usage:    "Demonstration of gorm with multiple databases",
		Action:   multipleDBAction,
		Category: "DB",
	})
}

func multipleDBAction(c *cli.Context) error {
	initDB1()
	initDB2()

	normalCrossDBQuery()
	sqlxCrossDBQuery()

	defer func() {
		_ = db1.Close()
		_ = db2.Close()
	}()

	return nil
}

func initDB1() {
	db1, _ = gorm.Open("sqlite3", db1Name)

	db1.SingularTable(true)
	db1.DB().SetMaxIdleConns(3)
	db1.LogMode(true)

	db1.AutoMigrate(&A{})

	var p A
	db1.Where(&A{
		Name: "p1",
	}).First(&p)
	if db1.HasTable(&A{}) && p.Name == "p1" {
		return
	}

	db1.Create(&A{
		ID:   id,
		Name: "p1",
	})
}

func initDB2() {
	db2, _ = gorm.Open("sqlite3", db2Name)

	db2.SingularTable(true)
	db2.DB().SetMaxIdleConns(3)
	db2.LogMode(true)

	db2.AutoMigrate(&B{})

	var p B
	db2.Where(&B{
		Type: "test type",
	}).First(&p)
	if db2.HasTable(&B{}) && p.Type == "test type" {
		return
	}

	db2.Create(&B{
		ID:   id,
		Type: "test type",
	})
}

func normalCrossDBQuery() {
	d1, err := sql.Open("sqlite3", db1Name)
	if err != nil {
		return
	}
	defer db1.Close()

	s := fmt.Sprintf("attach database '%s' as db2;", db2Name)
	r1, err := d1.Exec(s)
	fmt.Println(r1)

	s = fmt.Sprintf("select name, type from a a inner join db2.b b on b.ID = a.ID;")
	results, err := d1.Query(s)
	for results.Next() {
		var a, b interface{}
		if err := results.Scan(&a, &b); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		fmt.Println(a)
	}
	defer results.Close()

	r2, _ := d1.Exec("detach database db2;")
	fmt.Println(r2)
}

func sqlxCrossDBQuery() {
	d1, err := sqlx.Connect("sqlite3", db1Name)
	if err != nil {
		log.Fatal(err)
	}

	s := fmt.Sprintf("attach database '%s' as db2;", db2Name)
	r1, err := d1.Exec(s)
	fmt.Println(r1)

	s = fmt.Sprintf("select name, type from a a inner join db2.b b on b.ID = a.ID;")
	rows, err := d1.Queryx(s)
	for rows.Next() {
		var a, b interface{}
		err := rows.Scan(&a, &b)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v %#v\n", a, b)
	}
}
