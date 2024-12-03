package main

import (
	"log"
	"lumoshive-be-chap41/infra"
	"lumoshive-be-chap41/routes"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	_ "lumoshive-be-chap41/docs"
)

// @title Voucher API
// @version 1.0
// @description API for interacting with voucher
// @termsOfService http://example.com/terms/
// @contact.name Lumoshive Support
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email safiramadhani9@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatalf("can't init service context: %v", err)
	}

	r := routes.NewRoutes(*ctx)

	// Add Swagger endpoint
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
