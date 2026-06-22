package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"backend/config"
	"backend/model"
)

func GetRekomendasiWisata(w http.ResponseWriter, r *http.Request) {
	// Memberi tahu bahwa balasan kita berbentuk JSON
	w.Header().Set("Content-Type", "application/json")
    
	// Mengatur CORS agar aplikasi Frontend bisa mengambil data ini tanpa error
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Menangkap kata kunci dari URL, misalnya: ?kategori=Populer Tinggi
	kategori := r.URL.Query().Get("kategori")
	if kategori == "" {
		http.Error(w, `{"pesan": "Kategori belum dipilih. Contoh: ?kategori=Populer Tinggi"}`, http.StatusBadRequest)
		return
	}

	// Memerintahkan MySQL untuk mencari wisata sesuai kategori stiker yang diminta
	query := "SELECT id, title, categoryName, city, totalScore, reviewsCount, Cluster, Cluster_Name FROM data_wisata WHERE Cluster_Name = ?"
	rows, err := config.DB.Query(query, kategori)
	if err != nil {
		http.Error(w, `{"pesan": "Terjadi gangguan saat mengambil data di server"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Menampung hasil pencarian dari MySQL
	var daftarWisata []model.Wisata
	for rows.Next() {
		var wisata model.Wisata
		err := rows.Scan(&wisata.ID, &wisata.Title, &wisata.CategoryName, &wisata.City, &wisata.TotalScore, &wisata.ReviewsCount, &wisata.Cluster, &wisata.ClusterName)
		if err != nil {
            // BARIS BARU: Paksa Golang mencetak alasan errornya ke terminal
			fmt.Println("TERJADI ERROR BACA DATA:", err)
			continue
		}
		daftarWisata = append(daftarWisata, wisata)
	}

	// Membungkus balasannya agar terlihat rapi
	response := map[string]interface{}{
		"pesan": "Berhasil mengambil rekomendasi wisata",
		"total_data": len(daftarWisata),
		"data":  daftarWisata,
	}

	// Mengirim balasan ke pengguna
	json.NewEncoder(w).Encode(response)
}