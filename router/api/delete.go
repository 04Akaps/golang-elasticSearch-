package api

import (
	"elasticSearch/services"
	"elasticSearch/types"
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
	d := &Delete{
		engine:  engine,
		logger:  log15.New("router", "delete"),
		service: service,
	}

	baseUri := "delete"

	d.engine.POST(baseUri+"/delete-example/:name", d.deleteExample)
}

func (d *Delete) deleteExample(c *gin.Context) {
	var req types.DeleteUserReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = d.service.DeleteUser("user", req.Name); err != nil {
		errResponse(c, err.Error())
	} else {
		msg := fmt.Sprintf("Success To Delete Document -> Name : %s", req.Name)
		successResponse(c, msg)
	}
}
