package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Format koneksi: username:password@tcp(host:port)/nama_database
	// Sesuaikan jika MySQL-mu menggunakan password (biasanya bawaan XAMPP tidak pakai password)
	dsn := "root:@tcp(127.0.0.1:3306)/db_wisata"
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka pintu ke MySQL:", err)
	}

	// Mengecek apakah koneksinya benar-benar tersambung
	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon, pastikan MySQL/XAMPP sudah menyala:", err)
	}

	fmt.Println("Mantap! Koneksi ke database db_wisata berhasil!")
	DB = db
}