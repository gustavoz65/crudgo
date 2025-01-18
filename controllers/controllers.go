package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/api-go-gin/database"
	"github.com/gustavoz65/api-go-gin/models"
)

// função para carregar a nossa estrutura de aluno
func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// exemplo de JSON retornado
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ",Tudo Beleza?",
	})

}

// Função para criar um novo aluno no banco
func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// Função para buscar um aluno no banco
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id") //buscar ID como string

	//validar e converter o ID para interiro
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "ID não pode ser vazio"})
		return
	}

	//busca no banco de dados
	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// Função para deletar um aluno no banco
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "aluno deletado com sucesso"})
	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado"})
		return
	}
}

// Função para editar um aluno no banco
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// Função para buscar por CPF um aluno no banco
func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// Função para buscar por RG um aluno no banco
func BuscaAlunoPorRg(c *gin.Context) {
	var aluno models.Aluno
	rg := c.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
