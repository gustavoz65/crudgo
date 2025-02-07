package main

import (
	"github.com/gustavoz65/api-go-gin/database"
	"github.com/gustavoz65/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

//chamo a função para conectar ao banco e tambem a funçao para criar as rotas "EndPoints"
