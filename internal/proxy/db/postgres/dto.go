package proxy

import "time"

type ProxyDTO struct {
	ID       int
	Scheme   string
	Host     string
	Username string
	Password string
	Port     int
	CreateAt time.Time
	UpdateAt time.Time
	IsBad    bool
}
