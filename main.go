package main

import (
	"fmt"

	"github.com/CarlosGenuino/api-rest-go/database"
	"github.com/CarlosGenuino/api-rest-go/models"
	"github.com/CarlosGenuino/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Carmela Dutra", Historia: "Primeira Primeira dama do Brasil. Esposa do ex-presidente Dutra"},
		{Id: 2, Nome: "Person 2", Historia: "Ficticio"},
	}
	database.ConectaBancoDeDados()
	routes.HandleRequest()
	fmt.Println("executando servidor go!")
}
