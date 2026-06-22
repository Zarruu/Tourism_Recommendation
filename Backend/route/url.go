package route

import (
	"net/http"
	"backend/controller"
)

func SetupRoutes() {
	// Jika ada yang mengakses /api/rekomendasi, panggil fungsi GetRekomendasiWisata
	http.HandleFunc("/api/rekomendasi", controller.GetRekomendasiWisata)
}