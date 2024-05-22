package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/output"
)

type App struct {
	R *gin.Engine
}

func NewApp() *App {
	var app App
	r := gin.Default
	app.R = r()
	app.registerRoutes()
	return &app
}

func (a *App) registerRoutes() {
	a.R.GET("/ping", HandPing)
	a.R.GET("/health", HandHealth)
}

func HandPing(ctx *gin.Context) {
	id := ctx.Query("id")
	output.ReturnSuccessResponse(ctx, id)
}

func HandHealth(ctx *gin.Context) {
	output.ReturnSuccessResponse(ctx, "health~")
}
