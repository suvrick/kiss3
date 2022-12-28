package proxy

import (
	"context"
)

type IProxyService interface {
	Get(ctx context.Context, limit int) ([]Proxy, error)
	GetByID(ctx context.Context, proxyID uint64) (Proxy, error)
	Create(ctx context.Context, proxy Proxy) (Proxy, error)
	Update(ctx context.Context, proxy Proxy) (Proxy, error)
	Delete(ctx context.Context, proxyID uint64) error
}

type IProxyRepository interface {
	Get(ctx context.Context, limit int) ([]Proxy, error)
	GetByID(ctx context.Context, proxyID uint64) (Proxy, error)
	Create(ctx context.Context, proxy Proxy) (Proxy, error)
	Update(ctx context.Context, proxy Proxy) (Proxy, error)
	Delete(ctx context.Context, proxyID uint64) error
}
