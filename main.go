// bdTEST project main.go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var Dsn = os.Getenv("DSN")
	fmt.Println("А где блят, где?", Dsn)
	x, y, z := OpenDB(Dsn)

	fmt.Println("sql.DB:", x, "error:", y, "boolean:", z)
}

func OpenDB(dsn string) (*sql.DB, error, bool) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err, false
	}
	if err = db.Ping(); err != nil {
		return nil, err, false
	}
	return db, nil, true
}

/*




 */
