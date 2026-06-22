package main

import (
	"fmt"
	"net/http"
	"backend/config"
	"backend/route"
)

func main() {
	// 1. Nyalakan koneksi ke database dulu
	config.ConnectDB()

	// 2. Siapkan alamat-alamat URL-nya
	route.SetupRoutes()

	// 3. Nyalakan server Web Service
	port := ":8080"
	fmt.Println("Web Service sudah menyala! Silakan buka:")
	fmt.Println("http://localhost:8080/api/rekomendasi?kategori=Populer Tinggi")
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Gagal menyalakan server:", err)
	}
}