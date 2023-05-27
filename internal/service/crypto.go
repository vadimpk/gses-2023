package service

import "context"

type cryptoService struct {
	serviceContext
}

func NewCryptoService(opts *Options) *cryptoService {
	return &cryptoService{
		serviceContext: serviceContext{
			storages: opts.Storages,
			apis:     opts.APIs,
			logger:   opts.Logger.Named("CryptoService"),
			cfg:      opts.Cfg,
		},
	}
}

func (s *cryptoService) GetRate(ctx context.Context, opts *GetRateOptions) (int, error) {
	logger := s.logger.Named("GetRate").WithContext(ctx)

	logger.Info("successfully got rate")
	return 0, nil
}
