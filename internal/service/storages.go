package service

import "context"

type Storages struct {
	Email EmailStorage
}

type EmailStorage interface {
	Save(ctx context.Context, email string) error
	List(ctx context.Context) ([]string, error)
}
