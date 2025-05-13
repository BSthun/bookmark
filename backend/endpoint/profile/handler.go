package profileEndpoint

import (
	"bookmark-backend/common/config"
)

type Handler struct {
	config *config.Config
}

func Handle(config *config.Config) *Handler {
	h := &Handler{
		config: config,
	}

	return h
}
