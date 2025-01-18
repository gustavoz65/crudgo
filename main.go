package main

import (
	"github.com/gustavoz65/api-go-gin/models"
	"github.com/gustavoz65/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gustavo", CPF: "123.456.789-10", RG: "763928"},
		{Nome: "Maria", CPF: "167,877,942-10", RG: "337764"},
	}

	routes.HandleRequests()
}
