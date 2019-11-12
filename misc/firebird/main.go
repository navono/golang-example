package firebird

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:     "firebird",
		Aliases:  []string{"fb"},
		Usage:    "demonstration of firebird db",
		Action:   fbAgent,
		Category: "DB",
	})
}

func fbAgent(c *cli.Context) error {
	conn, err := sql.Open("firebirdsql",
		"supconadmin:supcon@localhost:3061/D:\\sourcecode\\go\\golang-example\\misc\\firebird\\TEST_DB.FDB")
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	var n int
	if err := conn.QueryRow("SELECT Count(*) FROM TAGS").Scan(&n); err != nil {
		return err
	}
	fmt.Println(n)

	sqlxDemo()

	return nil
}

func sqlxDemo() error {
	db, err := sqlx.Open("firebirdsql",
		"supconadmin:supcon@localhost:3061/D:\\sourcecode\\go\\golang-example\\misc\\firebird\\TEST_DB.FDB")
	if err != nil {
		return err
	}

	type tag struct {
		TagID   string `db:"TAG_ID"`
		TagName string `db:"TAG_NAME"`
	}
	//
	// var t tag
	// err = db.Get(&t, "SELECT TAG_ID, TAG_NAME FROM tags WHERE tag_name=$1", "XXX")

	var tl []tag
	err = db.Select(&tl, "SELECT TAG_ID, TAG_NAME FROM tags WHERE TAG_NAME='XXX'")

	if err != nil {
		return err
	}

	rows, err := db.NamedQuery(`SELECT TAG_ID, TAG_NAME FROM TAGS WHERE TAG_NAME=:tn`,
		map[string]interface{}{"tn": "XXX"})
	if err != nil {
		return err
	}
	for rows.Next() {
		// var a, b interface{}
		var a tag
		// var a interface{}
		_ = rows.StructScan(&a)

		fmt.Printf("%#v\n", a)
	}

	// rows, err := db.Query("SELECT Count(*) FROM TAGS")
	// for rows.Next() {
	// 	// var a, b interface{}
	// 	var a interface{}
	// 	_ = rows.Scan(&a)
	//
	// 	fmt.Printf("%#v\n", a)
	// }

	return nil
}
