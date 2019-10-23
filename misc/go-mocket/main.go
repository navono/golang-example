package go_mocket

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "mocket",
		Aliases: []string{"mocket"},

		Usage:    "Demonstration of go-mocket",
		Action:   mocketAction,
		Category: "DB",
	})
}

func mocketAction(c *cli.Context) error {

	return nil
}

func GetUsers(db *sql.DB) []map[string]string {
	var res []map[string]string
	age := 27
	rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var age string
		var id int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		row := map[string]string{"name": name, "age": age}
		res = append(res, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func GetUsersByGorm(db *gorm.DB) []map[string]string {
	age := 27
	type user struct {
		gorm.Model
		Name string
		Age  int
	}

	var res []map[string]string
	var u user
	db.Select("name").Where(&user{
		Age: age,
	}).Find(&u)

	res = append(res, map[string]string{
		"name": u.Name,
		"age":  fmt.Sprintf("%d", u.Age),
	})

	return res
}
