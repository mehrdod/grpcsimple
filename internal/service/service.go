package service

import (
	"github.com/mehrdod/grpcsimple/pkg/cache"
	"github.com/mehrdod/grpcsimple/pkg/logger"
)

type Fibonacci interface {
	Get(from, to int, fibChan chan int64) error
}

type Services struct {
	Fibonacci Fibonacci
}

func NewServices(cache cache.Cache, cacheLimit int, logger logger.Logger) *Services {
	return &Services{Fibonacci: NewFibonacciService(cache, cacheLimit, logger)}
}
