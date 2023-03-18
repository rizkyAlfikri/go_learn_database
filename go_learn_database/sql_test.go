package go_learn_database

import (
	"context"
	"fmt"
	"testing"
)

/*
Eksekusi Perintah SQL
- Saat membuat  aplikasi menggunkkana database, sudah pasti kita ingin berkomunikasi dengan datqabase mengunakan perintah SQL
  - Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function(DB) ExxecContext(context, sql, params)
  - Ketika mengirim perintah perintah SQL, kita bisa mengirim sinyal yang sudah pernah kita pelajari di course Golang Context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya
*/
func TestExeSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('budi', 'Budi')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name FROM customer "
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		t.Log(id)
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}
