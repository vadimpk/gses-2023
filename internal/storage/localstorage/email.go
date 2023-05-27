package localstorage

import (
	"context"
	"github.com/vadimpk/gses-2023/pkg/database"
)

type emailStorage struct {
	db *database.FileDB
}

func NewEmailStorage(db *database.FileDB) *emailStorage {
	return &emailStorage{
		db: db,
	}
}

const (
	emailStorageFileName = "emails.txt"
)

//Save(ctx context.Context, email string) error
//	List(ctx context.Context) ([]string, error)

func (s *emailStorage) Save(ctx context.Context, email string) error {
	return s.db.Append(ctx, emailStorageFileName, []byte(email))
}

func (s *emailStorage) List(ctx context.Context) ([]string, error) {
	data, err := s.db.Read(ctx, emailStorageFileName)
	if err != nil {
		return nil, err
	}
	// TODO: split data by new line
	return []string{string(data)}, nil
}
