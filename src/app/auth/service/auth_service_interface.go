package service

import (
	"context"
	"github.com/yaza-putu/online-test-dot/src/http/response"
)

type AuthInterface interface {
	Login(ctx context.Context, email string, password string) response.DataApi
	Refresh(ctx context.Context, oToken string) response.DataApi
}
