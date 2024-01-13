package handlers

import "github.com/gin-gonic/gin"

type (
	HandlerError struct {
		Error string `json:"error"`
	}
)

func HandleError(c *gin.Context, code int, err error) {
	e := HandlerError{
		Error: err.Error(),
	}
	c.AbortWithStatusJSON(code, e)
}
