package handler

import (
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

type (
	// Handler list handler
	Handler struct {
		DB    *mgo.Session
		REDIS *redis.Client
	}
)

const (
	Key = "secret"
)

func (h *Handler) getRedis() {
	return
}
