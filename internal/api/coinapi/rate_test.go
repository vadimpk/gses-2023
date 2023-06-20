package coinapi

import (
	"context"
	"testing"

	"github.com/vadimpk/gses-2023/internal/service"
	"github.com/vadimpk/gses-2023/pkg/logging"
)

func TestCoinAPI_GetRate(t *testing.T) {
	type args struct {
		opts *service.GetRateOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				opts: &service.GetRateOptions{
					CryptoCurrency: "BTC",
					Currency:       "UAH",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(&Options{
				ApiKey: "F9326003-515F-4655-A9A8-2ACF5D8E900F",
				Logger: logging.New("info"),
			})
			_, _ = c.GetRate(context.Background(), tt.args.opts)
			//assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
