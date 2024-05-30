// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/wandermaia/pos-golang-di/product"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	productRepository := product.NewProductRepository(db)
	productUseCase := product.NewProductUseCase(productRepository)
	return productUseCase
}

// wire.go:

var setRepositoryDependency = wire.NewSet(product.NewProductRepository, wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))
