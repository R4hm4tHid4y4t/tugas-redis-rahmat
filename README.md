# Tugas Topik Khusus - Implementasi Redis dengan Golang (Track Method)

Repositori ini berisi implementasi program sederhana menggunakan bahasa Go (Golang) dan Redis untuk memenuhi tugas mata kuliah Topik Khusus. Program ini mendemonstrasikan metode "Tracking" (misalnya untuk melacak jumlah klik atau *page view*) dengan memanfaatkan fitur `INCR` pada Redis.

## Prompt AI yang Digunakan
Penyelesaian tugas ini dibantu oleh AI dengan menggunakan prompt berikut:

> "suruh AI install redis beserta hal yang diperlukan menggunakan bahasa GOLANG pakai track method. prompt dibikin di readme.md. Tugas Topik Khusus: Install Redis, prompt sebuah program menggunakan bahasa pilihan seperti python, buat repository baru di github organization lalu push. Bantulah saya sampai tugas Topik Khusus saya selesai sesuai dangan pernyataan diatas"

## Cara Menjalankan Program

1. Pastikan Redis sudah terinstal dan berjalan di lokal (Port `6379`).
2. Pastikan Golang sudah terinstal.
3. Clone repositori ini ke komputer Anda.
4. Buka terminal dan arahkan ke dalam folder repositori ini.
5. Download *dependencies* yang dibutuhkan dengan perintah:
   ```bash
   go mod tidy