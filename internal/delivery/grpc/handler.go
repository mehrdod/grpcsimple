package grpcdelivery

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/mehrdod/grpcsimple/internal/service"
)

type Handler struct {
	services *service.Services
	UnimplementedFibonacciServer
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Get(fibRange *Range, stream Fibonacci_GetServer) error {
	if fibRange.From > fibRange.To {
		return errors.New("wrong params")
	}
	fibChan := make(chan int64)

	errs, _ := errgroup.WithContext(context.Background())
	errs.Go(func() error {
		return h.services.Fibonacci.Get(int(fibRange.From), int(fibRange.To), fibChan)
	})
	for fibNum := range fibChan {
		if err := stream.Send(&FibonacciNum{Value: fibNum}); err != nil {
			return err
		}
	}

	return errs.Wait()
}
