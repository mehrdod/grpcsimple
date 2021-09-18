package http

import (
	"io"
	"net/http"
	"strconv"

	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initFibonacci(api *gin.RouterGroup) {
	fibonacci := api.Group("/fibonacci")
	{
		fibonacci.GET("", h.fibonacciGet)
	}
}

type fibonacciRequest struct {
	From int `form:"from" binding:"required,min=1,max=100000"`
	To   int `form:"to" binding:"required,min=1,max=100000"`
}

func (h *Handler) fibonacciGet(c *gin.Context) {
	var fibParams fibonacciRequest
	err := c.ShouldBindQuery(&fibParams)

	if err != nil || fibParams.From > fibParams.To {
		h.newResponse(c, http.StatusBadRequest, ErrBadParams, err)
		return
	}
	errs, _ := errgroup.WithContext(c.Request.Context())
	fibChan := make(chan int64)

	errs.Go(func() error {
		return h.services.Fibonacci.Get(fibParams.From, fibParams.To, fibChan)
	})

	c.Stream(func(w io.Writer) bool {
		currentFib, ok := <-fibChan
		if !ok {
			return false
		}
		output := []byte(strconv.FormatInt(currentFib, 10))
		c.Writer.Write(append(output, []byte(", ")...))
		return true
	})
	err = errs.Wait()
	if err != nil {
		h.newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}
}
