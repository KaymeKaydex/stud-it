package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"test/internal/pkg/api/v1/user"
	"test/internal/pkg/service"
)

type Router struct {
	ctx context.Context

	engine *gin.Engine
	srv    *service.Service
}

func New(ctx context.Context, srv *service.Service) (*Router, error) {
	r := gin.Default()

	return &Router{
		ctx:    ctx,
		engine: r,
		srv:    srv,
	}, nil
}

func (r *Router) Run() error {
	gRouter := r.engine

	userGroup := gRouter.Group("/user")
	userController := user.NewController(r.srv)
	{
		userGroup.GET("/age", userController.GetAgeInfo)
	}

	gRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r.engine.Run("0.0.0.0:8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
