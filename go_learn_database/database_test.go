package go_learn_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3304)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	// gunakan database,
	// close database jika tidak digunakan
	defer db.Close()
}
