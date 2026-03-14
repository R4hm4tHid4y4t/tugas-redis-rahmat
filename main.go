package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	// Import module buatan kita sendiri
	"tugas-redis-rahmat/repository"
	"tugas-redis-rahmat/usecase"
)

func main() {
	// 1. Inisialisasi Koneksi Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}
	fmt.Println("==== Redis Terhubung dengan Baik! ====")

	// 2. Persiapan Dependency Injection (Struktur Clean Architecture)
	trackRepo := repository.NewRedisTrackRepository(client)
	trackUC := usecase.NewTrackUsecase(trackRepo)

	// 3. Menjalankan Aplikasi (Simulasi)
	fmt.Println("--- Memulai Aplikasi Tracking Event ---")

	// Simulasi event yang terjadi di aplikasi kita
	events := []string{
		"click_button_login",
		"click_button_login",
		"view_page_dashboard",
		"click_button_logout",
		"click_button_login",
	}

	// Looping untuk memanggil Track Method
	for _, event := range events {
		err := trackUC.RecordEvent(ctx, event)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}

	fmt.Println("\n==== Simulasi Selesai ====")
}
