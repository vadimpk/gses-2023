package service

import (
	"context"
	"fmt"
)

type emailService struct {
	serviceContext
	cryptoService CryptoService
}

func NewEmailService(opts *Options, cryptoService CryptoService) *emailService {
	return &emailService{
		serviceContext: serviceContext{
			storages: opts.Storages,
			apis:     opts.APIs,
			logger:   opts.Logger.Named("EmailService"),
			cfg:      opts.Cfg,
		},
		cryptoService: cryptoService,
	}
}

func (s *emailService) Subscribe(ctx context.Context, email string) error {
	logger := s.logger.Named("Subscribe").
		WithContext(ctx).
		With("email", email)

	existingEmail, err := s.storages.Email.Get(ctx, email)
	if err != nil {
		logger.Error("failed to get email from storage", "err", err)
		return fmt.Errorf("failed to get email from storage: %w", err)
	}
	if existingEmail != "" {
		logger.Info("email already exists")
		return ErrSubscribeAlreadySubscribed
	}

	err = s.storages.Email.Save(ctx, email)
	if err != nil {
		logger.Error("failed to save email", "err", err)
		return fmt.Errorf("failed to save email to storage: %w", err)
	}

	logger.Info("successfully subscribed")
	return nil
}

func (s *emailService) SendRateInfo(ctx context.Context) error {
	logger := s.logger.Named("SendRateInfo").
		WithContext(ctx)

	logger.Info("successfully sent rate info")
	return nil
}
