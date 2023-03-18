package go_learn_database

import (
	"database/sql"
	"time"
)

/*
	Database pooling
	- sql.DB di Golang sebenarnya bukanlah koneksi ke database
	- Melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
	- Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan tidak perlu melakukan management koneksi database secara manual
	- Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksumal koneksi yang dibuat oleh Golang, sehingga tidak membanjiri koneksi ke database, karena boasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita bisa gunakan
*/

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // Pengaturan berapa jumlah koneksi minimal yang dibuat
	db.SetMaxOpenConns(100)                 // Pengaturan berapa jumlah koneksi maksimal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // Pengaturan berapa lama koneksi yang tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // Pengaturan berapa lama koneksi yang boleh digunakan

	return db
}
