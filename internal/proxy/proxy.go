package proxy

import (
	"time"
)

type Proxy struct {
	ID       int       `json:"id,omitempty"`
	Scheme   string    `json:"scheme,omitempty"`
	Host     string    `json:"host,omitempty"`
	Port     int       `json:"port,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	AtCreate time.Time `json:"create_at,omitempty"`
	AtUpdate time.Time `json:"update_at,omitempty"`
	IsBad    bool      `json:"is_bad,omitempty"`
}
