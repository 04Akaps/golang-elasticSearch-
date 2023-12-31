package api

import (
	"elasticSearch/services"
	"elasticSearch/types"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Search struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Search
}

func NewSearchApi(engine *gin.Engine, service *services.Search) {
	search := &Search{
		engine:  engine,
		logger:  log15.New("router", "search"),
		service: service,
	}

	baseUri := "search"

	search.engine.POST("/all", search.findAll)
	search.engine.POST(baseUri+"/user-by-name/:name", search.searchUserByName)
	search.engine.POST(baseUri+"/user-by-age/:age", search.searchUserByAge)
	search.engine.POST(baseUri+"/user-by-address/:address", search.searchUserByAddress)

}

func (m *Search) findAll(c *gin.Context) {
	var req types.FindAllReq

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.FindAllByV8(types.Index, req.Size); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}

func (m *Search) searchUserByName(c *gin.Context) {
	var req types.SearchByNameReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByName(types.Index, req.Name, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}

func (m *Search) searchUserByAge(c *gin.Context) {
	var req types.SearchByAgeReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByAge(types.Index, req.Age, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}

func (m *Search) searchUserByAddress(c *gin.Context) {
	var req types.SearchByAddressReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByAddress(types.Index, req.Address, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}
