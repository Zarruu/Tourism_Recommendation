export default function EvaluasiModel() {
  return (
    <div className="evaluasi-container">
      <h2>Hasil Evaluasi K-Means Clustering</h2>
      <p>Data tempat wisata telah berhasil dikelompokkan menjadi 3 klaster menggunakan algoritma K-Means.</p>
      
      <div className="metrik-box">
        <div className="metrik">
          <h3>Silhouette Score</h3>
          <p className="nilai">0.488</p>
          <small>Mendekati 1 berarti kelompok sudah sangat baik dan tidak tumpang tindih.</small>
        </div>
        <div className="metrik">
          <h3>Inertia (WCSS)</h3>
          <p className="nilai">192.350</p>
          <small>Semakin kecil semakin baik tingkat kerapatan di dalam kelompoknya.</small>
        </div>
      </div>

      <div className="visualisasi">
        <h3>Visualisasi Persebaran Klaster</h3>
        {/* Pastikan kamu memindahkan gambar hasil grafik Python ke folder public/ */}
        <img src="/kmeans_jabar.png" alt="Grafik K-Means Jawa Barat" style={{maxWidth: '100%', borderRadius: '8px'}} />
      </div>
    </div>
  );
}