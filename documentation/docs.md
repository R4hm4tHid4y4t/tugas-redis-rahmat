# Dokumentasi Program: Redis Track Method (Clean Architecture)

Dokumen ini menjelaskan alur kerja dan struktur dari program pelacakan aktivitas (Track Method) yang diimplementasikan menggunakan Golang dan Redis.

## Struktur Direktori

Program ini memisahkan logika ke dalam beberapa layer (Clean Architecture) agar kode lebih rapi, mudah dibaca, dan mudah di-*maintenance*:

1. **`domain/track.go`**
   - Merupakan layer pusat yang berisi *Interface* (kontrak).
   - Menyimpan definisi `TrackRepository` (untuk interaksi database) dan `TrackUsecase` (untuk logika bisnis).

2. **`repository/redis_repository.go`**
   - Layer yang bertanggung jawab langsung berkomunikasi dengan Redis.
   - Menggunakan perintah `INCR` pada Redis. Setiap kali fungsi dipanggil dengan nama *event* tertentu, Redis akan menambah nilainya sebanyak 1.

3. **`usecase/track_usecase.go`**
   - Layer yang memproses logika bisnis.
   - Mengatur pemanggilan ke layer repository dan memberikan respon (berupa log/cetakan ke terminal) jika aktivitas berhasil dicatat.

4. **`main.go`**
   - Sebagai titik awal aplikasi berjalan.
   - Melakukan inisialisasi koneksi ke server Redis lokal (Port 6379).
   - Menyatukan ( *Dependency Injection*) antara Repository dan Usecase.
   - Menjalankan simulasi beberapa aktivitas seperti `click_button_login` dan `view_page_dashboard`.