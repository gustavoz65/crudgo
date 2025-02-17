package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/gustavoz65/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos", controller.ExibeAlunos)
	r.POST("/alunos", controller.CriarNovoAluno)
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCpf)
	r.GET("/alunos/rg/:rg", controller.BuscaAlunoPorRg)
	r.GET("/index", controller.ExibePagina)
	r.NoRoute(controller.RotaNaoEncontrada)
	r.Run()
}
   // requisição ==== endPoint =====  função que vai ser executada