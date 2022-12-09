package exception

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/model"
	"net/http"
)

func ErrorHandler(c *gin.Context, err any) {
	unathorizedError, ok := err.(UnauthorizedError)
	if ok {
		AbortWithStatus(c, http.StatusUnauthorized, unathorizedError)
		return
	}

	errr, ok := err.(error)
	if !ok {
		errr = errors.New("unknown error")
	}

	AbortWithStatus(c, http.StatusInternalServerError, errr)
}

func AbortWithStatus(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, model.WebResponse{
		Code:    code,
		Message: err.Error(),
	})
}
