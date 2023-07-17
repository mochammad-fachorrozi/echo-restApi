package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Ganti informasi koneksi database dengan yang sesuai
	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "echo_restApi"

	// Format string koneksi database
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuka koneksi ke database MySQL
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	// Ping database untuk memastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db

	fmt.Println("Koneksi ke database berhasil")
}
