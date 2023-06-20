package service

import (
	"context"
	"testing"

	"github.com/vadimpk/gses-2023/internal/storage/localstorage"
	"github.com/vadimpk/gses-2023/pkg/database"
	"github.com/vadimpk/gses-2023/pkg/logging"
)

func TestEmailService_Subscribe(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "test",
			email:   "test",
			wantErr: true,
		},
		{
			name:    "test",
			email:   "test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &emailService{
				serviceContext: serviceContext{
					storages: Storages{
						localstorage.NewEmailStorage(database.NewFileDB("local/")),
					},
					logger: logging.New("info"),
				},
			}
			if err := s.Subscribe(context.Background(), tt.email); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
