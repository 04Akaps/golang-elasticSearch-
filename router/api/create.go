package api

import (
	"elasticSearch/services"
	"elasticSearch/types"
	"elasticSearch/types/schema"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Create struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Create
}

func NewCreateApi(engine *gin.Engine, service *services.Create) {
	create := &Create{
		engine:  engine,
		logger:  log15.New("router", "create"),
		service: service,
	}

	baseUri := "create"

	create.engine.POST(baseUri+"/user", create.createUser)

}

func (m *Create) createUser(c *gin.Context) {
	var req types.CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
		return
	}

	if err := m.service.CreateData(types.Index, schema.User{
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	}); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, "success add new User")
	}

}
