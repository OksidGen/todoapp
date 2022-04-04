package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OksidGen/todoapp/internal/middleware"
)

func InitializeRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("./web/templates/*.html")
	router.Static("/www", "./web/static/")
	router.StaticFile("/favicon.ico", "./web/favicon.ico")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", middleware.GetAllTasks)
	router.POST("/create", middleware.CreateTask)
	router.POST("/update", middleware.UpdateTask)
	router.POST("/status", middleware.UpdateStatusTask)
	router.POST("/delete", middleware.DeleteTask)
	router.POST("/deleteall", middleware.DeleteAllTask)

	router.Run()
}
