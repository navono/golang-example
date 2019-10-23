package go_mocket

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

var (
	DB  *sql.DB
	GDB *gorm.DB
)

func SetupSQLTests() *sql.DB { // or *gorm.DB
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true

	// Regular sql package usage
	db, err := sql.Open(mocket.DriverName, "connection_string")
	if err != nil {
		return nil
	}

	return db
}

func SetupGormTests() *gorm.DB {
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true

	// GORM
	db, err := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string
	if err != nil {
		return nil
	}

	return db
}

func TestResponses(t *testing.T) {
	DB = SetupSQLTests()
	GDB = SetupGormTests()

	t.Run("Simple SELECT caught by query", func(t *testing.T) {
		mocket.Catcher.Logging = true
		// Important: Use database files here (snake_case) and not struct variables (CamelCase)
		// eg: first_name, last_name, date_of_birth NOT FirstName, LastName or DateOfBirth
		commonReply := []map[string]interface{}{{"user_id": 27, "name": "FirstLast", "age": "30"}}
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT name FROM users WHERE`).WithReply(commonReply)

		normalResult := GetUsers(DB) // Global or local variable
		if len(normalResult) != 1 {
			t.Errorf("Returned sets is not equal to 1. Received %d", len(normalResult))
		}
		if normalResult[0]["age"] != "30" {
			t.Errorf("Age is not equal. Got %v", normalResult[0]["age"])
		}

		gormResult := GetUsersByGorm(GDB) // Global or local variable
		if len(gormResult) != 1 {
			t.Errorf("Returned sets is not equal to 1. Received %d", len(gormResult))
		}
		if gormResult[0]["age"] != "30" {
			t.Errorf("Age is not equal. Got %v", gormResult[0]["age"])
		}
	})
}
