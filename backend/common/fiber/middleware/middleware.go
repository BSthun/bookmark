package middleware

import "bookmark-backend/common/config"

type Middleware struct {
	config *config.Config
}

func Init(config *config.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}
