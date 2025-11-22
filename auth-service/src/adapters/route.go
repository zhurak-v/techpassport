package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
	})
}

type RouterSetup struct {
	Engine *gin.Engine
}

func NewRouterSetup(engine *gin.Engine) *RouterSetup {
	return &RouterSetup{
		Engine: engine,
	}
}

func (r *RouterSetup) SetupRoutes() {
	r.Engine.GET("/hello", HelloWorldHandler)
}
