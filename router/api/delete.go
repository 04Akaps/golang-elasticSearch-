package api

import (
	"elasticSearch/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Delete struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Delete
}

func NewDeleteApi(engine *gin.Engine, service *services.Delete) {
	delete := &Delete{
		engine:  engine,
		logger:  log15.New("router", "delete"),
		service: service,
	}

	fmt.Println(delete)
}
