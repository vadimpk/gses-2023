package service

import "context"

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
	logger := s.logger.Named("Subscribe").WithContext(ctx)

	logger.Info("successfully subscribed")
	return nil
}

func (s *emailService) SendRateInfo(ctx context.Context) error {
	logger := s.logger.Named("SendRateInfo").WithContext(ctx)

	logger.Info("successfully sent rate info")
	return nil
}
