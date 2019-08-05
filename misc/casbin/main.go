package casbin

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbPolicy()
}

func dbPolicy() {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)", false)
	if err != nil {
		fmt.Printf("LoadPolicy failed %v\n", err)
		return
	}

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e := casbin.NewEnforcer("./misc/casbin/rbac_model.conf", a)
	e.EnableLog(true)

	// Load the policy from DB.
	if err := e.LoadPolicy(); err != nil {
		fmt.Printf("LoadPolicy failed %v\n", err)
		return
	}

	e.AddPolicy("alice", "data1", "read")

	// Check the permission.
	if e.Enforce("alice", "data1", "read") == true {
		fmt.Println("permit")
	}

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	if err := e.SavePolicy(); err != nil {
		fmt.Printf("SavePolicy failed %v\n", err)
	}
}

func localPolycy() {
	e := casbin.NewEnforcer("./misc/casbin/rbac_model.conf", "./misc/casbin/rbac_policy.csv")
	e.EnableLog(true)

	//e2 := casbin.NewSyncedEnforcer("./misc/casbin/rbac_model.conf", "./misc/casbin/rbac_policy.csv")
	//e2.StartAutoLoadPolicy(5 * time.Second)

	sub := "abc123" // 想要访问资源的用户
	obj := "user"   // 将要访问的资源
	act := "GET"    // 用户在资源上执行的操作
	if e.Enforce(sub, obj, act) == true {
		// 允许操作
		fmt.Println("permit")
	} else {
		// 不允许操作
		fmt.Println("deny")
	}

	//roles := e.GetAllRoles()
	//fmt.Printf("All roles: %v\n", roles)

	//subs := e.GetAllSubjects()
	//fmt.Printf("All subjects: %v\n", subs)

	//uRoles := e.GetImplicitRolesForUser("super_admin")
	//fmt.Printf("All subjects: %v\n", uRoles)
}
