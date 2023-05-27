package service

import "context"

type APIs struct {
	Email  EmailAPI
	Crypto CryptoAPI
}

// EmailAPI provides methods for sending emails that are used in EmailService and
// implemented in external packages.
type EmailAPI interface {
	Send(ctx context.Context, opts *SendOptions) error
}

type SendOptions struct {
	To    string
	Title string
	Body  string
}

// CryptoAPI provides methods for getting crypto rates that are used in CryptoService and
// implemented in external packages.
type CryptoAPI interface {
	GetRate(ctx context.Context, opts *GetRateOptions) (int, error)
}
