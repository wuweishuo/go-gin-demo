package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/pkg/setting"
	v1 "go-gin-demo/routers/api/v1"
)

func InitRouters() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	tags := engine.Group("/api/v1/tags")
	{
		tags.GET("/", v1.GetTags)
		tags.POST("/", v1.AddTags)
		tags.PUT("/:id", v1.EditTags)
		tags.DELETE("/:id", v1.DeleteTags)
	}

	return engine
}
