package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// // Criando um  novo repositório
	// repository := product.NewProductRepository(db)

	// // Criando um novo usecase de produto
	// usecase := product.NewProductUseCase(repository)

	// As dependências foram substituídas pela função gerada no arquivo wire_gen.go
	usecase := NewUseCase(db)

	// utilizando o usecase para executear get product
	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)

}
