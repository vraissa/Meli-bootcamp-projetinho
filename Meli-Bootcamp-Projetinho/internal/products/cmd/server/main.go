package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meli-bootcamp/cmd/server/handler"
	"github.com/meli-bootcamp/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
