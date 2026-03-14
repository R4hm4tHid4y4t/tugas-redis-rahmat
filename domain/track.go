package domain

import "context"

// TrackRepository adalah kontrak untuk interaksi dengan database (Redis)
type TrackRepository interface {
	IncrementEvent(ctx context.Context, eventName string) (int64, error)
}

// TrackUsecase adalah kontrak untuk logika bisnis aplikasi
type TrackUsecase interface {
	RecordEvent(ctx context.Context, eventName string) error
}
