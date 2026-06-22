import { useState, useEffect } from 'react';

export default function RekomendasiWisata() {
  const [kategori, setKategori] = useState('Populer Tinggi');
  const [daftarWisata, setDaftarWisata] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    fetch(`http://localhost:8080/api/rekomendasi?kategori=${encodeURIComponent(kategori)}`)
      .then((response) => response.json())
      .then((data) => {
        setDaftarWisata(data.data || []);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Gagal mengambil data:", error);
        setLoading(false);
      });
  }, [kategori]);

  return (
    <div className="rekomendasi-wrapper">
      {/* Hero Section */}
      <div className="hero-section">
        <h2>Eksplorasi Keindahan Jawa Barat</h2>
        <p>Temukan destinasi wisata yang paling pas dengan preferensi liburanmu. Dari yang paling viral hingga surga tersembunyi.</p>
        
        <div className="filter-modern">
          <label>Pilih Suasana Liburanmu:</label>
          <select value={kategori} onChange={(e) => setKategori(e.target.value)}>
            <option value="Populer Tinggi">🌟 Super Viral & Ramai (Populer Tinggi)</option>
            <option value="Populer Sedang">🏖️ Santai & Aman (Populer Sedang)</option>
            <option value="Cluster Khusus/Niche">🍃 Hidden Gem & Spesifik (Niche)</option>
          </select>
        </div>
      </div>

      {/* Hasil Rekomendasi */}
      <div className="hasil-section">
        <h3 className="section-title">
          Rekomendasi untuk "{kategori === 'Cluster Khusus/Niche' ? 'Hidden Gem' : kategori}"
        </h3>
        
        {loading ? (
          <div className="loading-spinner">Mencarikan destinasi terbaik...</div>
        ) : daftarWisata.length === 0 ? (
          <div className="empty-state">Belum ada data wisata untuk kategori ini.</div>
        ) : (
          <div className="wisata-grid">
            {daftarWisata.map((wisata) => (
              <div key={wisata.id} className="wisata-card">
                <div className="card-image-placeholder">
                  <span className="badge-kota">{wisata.city}</span>
                </div>
                <div className="card-content">
                  <span className="category-tag">{wisata.categoryName}</span>
                  <h4>{wisata.title}</h4>
                  <div className="stats">
                    <span className="rating">⭐ {wisata.totalScore}/5.0</span>
                    <span className="reviews">💬 {wisata.reviewsCount} Ulasan</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}