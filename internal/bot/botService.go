package bot

import (
	"context"
	"time"
)

type BotService struct {
	repo IBotRepository
}

func NewBotService(repo IBotRepository) *BotService {
	return &BotService{
		repo: repo,
	}
}

func (srv *BotService) Get(ctx context.Context, limit int) ([]Bot, error) {
	return srv.repo.Get(ctx, limit)
}

func (srv *BotService) GetByID(ctx context.Context, botID uint64) (Bot, error) {
	return srv.repo.GetByID(ctx, botID)
}

func (srv *BotService) Create(ctx context.Context, bot Bot) (Bot, error) {
	bot.AtCreate = time.Now()
	return srv.repo.Create(ctx, bot)
}

func (srv *BotService) Update(ctx context.Context, bot Bot) (Bot, error) {
	bot.AtUpdate = time.Now()
	return srv.repo.Update(ctx, bot)
}

func (srv *BotService) Delete(ctx context.Context, botID uint64) error {
	return srv.repo.Delete(ctx, botID)
}
