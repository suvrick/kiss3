package proxy

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suvrick/kiss/internal/middlewares"
	"github.com/suvrick/kiss/internal/until"
)

type ProxyController struct {
	service IProxyService
}

func NewProxyController(service IProxyService) *ProxyController {
	return &ProxyController{
		service: service,
	}
}

func (ctrl *ProxyController) Register(router *gin.Engine) {
	proxyGroup := router.Group("api/proxies")

	proxyGroup.Use(middlewares.AuthMiddleware())

	proxyGroup.GET("/", ctrl.Get)
	proxyGroup.PUT("/", ctrl.Update)
	proxyGroup.POST("/", ctrl.Create)
	proxyGroup.DELETE("/:id", ctrl.Delete)
}

// curl -X GET localhost:8080/api/proxies
func (ctrl *ProxyController) Get(c *gin.Context) {

	proxies, err := ctrl.service.Get(c.Request.Context(), 0)

	if err != nil {
		log.Printf("[Get] Get proxy list fail. %v\n", err)
		until.HTTPResponse(c, http.StatusBadRequest, "Get proxy list fail", err, nil)
		return
	}

	log.Printf("[Get] Get proxy list successed. Count proxies: %d\n", len(proxies))
	until.HTTPResponse(c, http.StatusOK, "Get proxy list success", nil, proxies)
}

// curl -X POST localhost:8080/api/proxies/ --data {\"proxy\":\"qwe\"}
func (ctrl *ProxyController) Create(c *gin.Context) {

	dto := Proxy{}

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("[Create] Create proxy fail. %v\n", err)
		until.HTTPResponse(c, http.StatusBadRequest, "Create proxy fail.", nil, nil)
		return
	}

	p, err := ctrl.service.Create(c.Request.Context(), dto)
	if err != nil {
		log.Printf("[Create] Create proxy fail. %v, %v\n", p, err)
		until.HTTPResponse(c, http.StatusBadRequest, "Create proxy fail.", nil, nil)
		return
	}

	log.Printf("[Create] Create proxy successed. %v\n", p)
	until.HTTPResponse(c, http.StatusOK, "Create proxy success!", nil, p)
}

// curl -X PUT localhost:8080/api/proxies/ --data {\"id\":2,\"is_bad\":true}
func (ctrl *ProxyController) Update(c *gin.Context) {

	dto := Proxy{}

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("[Update] Update proxy fail.%v, %v\n", dto, err)
		until.HTTPResponse(c, http.StatusBadRequest, "Update proxy fail.", nil, nil)
		return
	}

	dto, err = ctrl.service.Update(c.Request.Context(), dto)
	if err != nil {
		log.Printf("[Update] Update proxy fail. %v, %v\n", dto, err)
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", nil, nil)
		return
	}

	log.Printf("[Update] Update proxy successed! %v\n", dto)
	until.HTTPResponse(c, http.StatusOK, "Update proxy successed!", nil, dto)
}

// curl -X DELETE localhost:8080/api/proxies/1
func (ctrl *ProxyController) Delete(c *gin.Context) {

	ID := c.Param("id")
	proxyID, err := strconv.Atoi(ID)
	if err != nil {
		log.Printf("[Delete] Delete proxy fail. %v, %v\n", proxyID, err)
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", nil, nil)
		return
	}

	err = ctrl.service.Delete(c.Request.Context(), uint64(proxyID))
	if err != nil {
		log.Printf("[Delete] Delete proxy fail. %v, %v\n", proxyID, err)
		until.HTTPResponse(c, http.StatusBadRequest, "Delete proxy fail.", nil, nil)
		return
	}

	log.Printf("[Delete] Delete proxy successed. %v\n", proxyID)
	until.HTTPResponse(c, http.StatusOK, "Delete proxy successed!", nil, nil)
}
