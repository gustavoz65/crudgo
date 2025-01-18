package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/gustavoz65/api-go-gin/controllers"
)


func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controller.ExibeAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.Run()
}
