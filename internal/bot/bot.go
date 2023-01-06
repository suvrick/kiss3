package bot

import (
	"time"
)

type BotLogin struct {
}

type Bot struct {
	ID       uint64    `json:"id,omitempty"`
	OwerID   uint64    `json:"ower_id,omitempty"`
	GameID   uint32    `json:"game_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Balance  uint32    `json:"balance,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Profile  string    `json:"profile,omitempty"`
	AtCreate time.Time `json:"at_create,omitempty"`
	AtUpdate time.Time `json:"at_update,omitempty"`
}
