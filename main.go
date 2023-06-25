package main

import (
	"github.com/krogertechnology/krogo/pkg/krogo"

	productsHandler "practice-app/handler/products"
	variantsHandler "practice-app/handler/variants"
	productsService "practice-app/service/products"
	variantsService "practice-app/service/variants"
	productsStore "practice-app/store/products"
	variantsStore "practice-app/store/variants"
)

func main() {
	app := krogo.New()
	app.Server.ValidateHeaders = false

	variantStore := variantsStore.New()
	productStore := productsStore.New(variantStore)

	productService := productsService.New(productStore, variantStore)
	variantService := variantsService.New(variantStore)

	productHandler := productsHandler.New(productService)
	variantHandler := variantsHandler.New(variantService)

	app.GET("/products/{id}", productHandler.GetByID)
	app.GET("/products", productHandler.GetAll)
	app.POST("/products", productHandler.Create)

	app.GET("/products/{pid}/variant/{id}", variantHandler.GetByID)
	app.POST("/products/{pid}/variant", variantHandler.Create)

	app.Start()
}
