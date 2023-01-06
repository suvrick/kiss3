package proxy

import (
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

	limit := c.GetInt("limit")

	proxies, err := ctrl.service.Get(c.Request.Context(), limit)

	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Get proxy list fail", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Get proxy list success", nil, proxies)
}

// curl -X POST localhost:8080/api/proxies/ --data {\"proxy\":\"qwe\"}
func (ctrl *ProxyController) Create(c *gin.Context) {

	// type ProxyDTOCreate struct {
	// 	Proxy     string `json:"proxy"`
	// 	Delimiter string `json:"delimiter,omitempty"`
	// }

	dto := Proxy{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	p, err := ctrl.service.Create(c.Request.Context(), dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Create success!", nil, p)
}

// curl -X PUT localhost:8080/api/proxies/ --data {\"id\":2,\"is_bad\":true}
func (ctrl *ProxyController) Update(c *gin.Context) {

	dto := Proxy{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Update fail.", err, nil)
		return
	}

	dto, err = ctrl.service.Update(c.Request.Context(), dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Create success!", nil, dto)
}

// curl -X DELETE localhost:8080/api/proxies/1
func (ctrl *ProxyController) Delete(c *gin.Context) {

	ID := c.Param("id")
	proxyID, err := strconv.Atoi(ID)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	err = ctrl.service.Delete(c.Request.Context(), uint64(proxyID))
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Delete success!", nil, nil)
}
