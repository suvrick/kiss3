package proxy

import (
	"context"
	"time"
)

type ProxyService struct {
	repo IProxyRepository
}

func NewProxyService(repo IProxyRepository) *ProxyService {
	return &ProxyService{
		repo: repo,
	}
}

func (srv *ProxyService) Get(ctx context.Context, limit int) ([]Proxy, error) {
	return srv.repo.Get(ctx, limit)
}

func (srv *ProxyService) GetByID(ctx context.Context, proxyID uint64) (Proxy, error) {
	return srv.repo.GetByID(ctx, proxyID)
}

func (srv *ProxyService) Create(ctx context.Context, proxy Proxy) (Proxy, error) {
	proxy.AtCreate = time.Now()
	return srv.repo.Create(ctx, proxy)
}

func (srv *ProxyService) Update(ctx context.Context, proxy Proxy) (Proxy, error) {
	proxy.AtUpdate = time.Now()
	return srv.repo.Update(ctx, proxy)
}

func (srv *ProxyService) Delete(ctx context.Context, proxyID uint64) error {
	return srv.repo.Delete(ctx, proxyID)
}
