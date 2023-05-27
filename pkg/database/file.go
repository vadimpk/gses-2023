package database

import "context"

type FileDB struct {
	baseFilePath string
}

func NewFileDB(baseFilePath string) *FileDB {
	return &FileDB{
		baseFilePath: baseFilePath,
	}
}

func (f *FileDB) Close() error {
	return nil
}

// Ping - checks if base file path exists.
func (f *FileDB) Ping(ctx context.Context) error {

	return nil
}

// Append - appends data to file.
func (f *FileDB) Append(ctx context.Context, file string, data []byte) error {

	return nil
}

// Read - reads data from file.
func (f *FileDB) Read(ctx context.Context, file string) ([]byte, error) {

	return nil, nil
}
