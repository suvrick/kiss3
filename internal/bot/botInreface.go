package bot

import (
	"context"
)

type IBotService interface {
	Get(ctx context.Context, limit int) ([]Bot, error)
	GetByID(ctx context.Context, botID uint64) (Bot, error)
	Create(ctx context.Context, bot Bot) (Bot, error)
	Update(ctx context.Context, bot Bot) (Bot, error)
	Delete(ctx context.Context, botID uint64) error
}

type IBotRepository interface {
	Get(ctx context.Context, limit int) ([]Bot, error)
	GetByID(ctx context.Context, botID uint64) (Bot, error)
	Create(ctx context.Context, bot Bot) (Bot, error)
	Update(ctx context.Context, bot Bot) (Bot, error)
	Delete(ctx context.Context, botID uint64) error
}
