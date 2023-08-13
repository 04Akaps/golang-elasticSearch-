package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func errResponse(c *gin.Context, response string) {
	c.JSON(http.StatusBadRequest, response)
}

func successResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}
