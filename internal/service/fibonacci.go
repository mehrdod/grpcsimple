package service

import (
	"strconv"

	"github.com/mehrdod/grpcsimple/pkg/cache"
	"github.com/mehrdod/grpcsimple/pkg/logger"
)

type FibonacciService struct {
	logger     logger.Logger
	cache      cache.Cache
	cacheLimit int
}

func NewFibonacciService(cache cache.Cache, cacheLimit int, logger logger.Logger) *FibonacciService {
	return &FibonacciService{
		cache:      cache,
		cacheLimit: cacheLimit,
		logger:     logger,
	}
}

func (f *FibonacciService) Get(from, to int, fibChan chan int64) error {

	maxFib, err := f.cache.Get("max")
	if err != nil || maxFib <= 2 {
		//nothing in cache
		err = f.calculateWithoutCache(1, 1, 1, to)
		if err != nil {
			close(fibChan)
			return err
		}
	} else if int(maxFib) < to {
		//just ignoring errors, but I know I shouldn't
		current, _ := f.cache.Get(strconv.Itoa(int(maxFib) - 1))
		next, _ := f.cache.Get(strconv.Itoa(int(maxFib)))
		err = f.calculateWithoutCache(int(maxFib)-1, current, next, to)
		if err != nil {
			close(fibChan)
			return err
		}
	}

	err = f.getAndSendFromCache(from, to, fibChan)
	if err != nil {
		f.logger.Errorf("Something wrong with caching: %s", err)
		//calculate it manually
	}
	return nil
}

func (f *FibonacciService) getAndSendFromCache(from, to int, fibChan chan int64) error {
	for i := from; i <= to; i++ {
		fibNum, err := f.cache.Get(strconv.Itoa(i))

		if err != nil {
			return err
		}
		fibChan <- fibNum
	}
	close(fibChan)
	return nil
}

func (f *FibonacciService) calculateWithoutCache(currentIndx int, current, next int64, to int) error {
	for i := currentIndx; i <= to; i++ {
		err := f.cache.Set(strconv.Itoa(i), current)
		if err != nil {
			return err
		}
		current, next = next, current+next
	}
	f.cache.Set("max", int64(to))
	return nil
}
