package usecase

import (
	"context"
	"fmt"

	"tugas-redis-rahmat/domain"
)

type trackUsecase struct {
	repo domain.TrackRepository
}

// NewTrackUsecase adalah fungsi pembuat (constructor)
func NewTrackUsecase(repo domain.TrackRepository) domain.TrackUsecase {
	return &trackUsecase{
		repo: repo,
	}
}

// RecordEvent menjalankan logika bisnis untuk mencatat aktivitas
func (u *trackUsecase) RecordEvent(ctx context.Context, eventName string) error {
	count, err := u.repo.IncrementEvent(ctx, eventName)
	if err != nil {
		return fmt.Errorf("gagal mencatat event %s: %w", eventName, err)
	}

	fmt.Printf("[TRACKING LOG] Aktivitas '%s' berhasil dicatat. Total kunjungan/klik saat ini: %d\n", eventName, count)
	return nil
}
