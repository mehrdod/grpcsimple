package http

import "github.com/gin-gonic/gin"

type dataResponse struct {
	Data interface{} `json:"data"`
}

type errResponse struct {
	Message string `json:"message"`
}

func newErrResponse(message string) *errResponse {
	return &errResponse{Message: message}
}

var (
	ErrBadParams      = newErrResponse("wrong params")
	ErrInternalServer = newErrResponse("oops, something went wrong")
)

func (h *Handler) newResponse(c *gin.Context, statusCode int, resp *errResponse, err error) {
	message := err.Error()
	h.logger.Errorf("%s", message)
	c.AbortWithStatusJSON(statusCode, resp)
}
