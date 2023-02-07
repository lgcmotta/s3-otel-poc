package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/lgcmotta/s3-poc/docs"
	"github.com/lgcmotta/s3-poc/endpoints"
	"github.com/lgcmotta/s3-poc/otel"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// @title AWS S3 OTEL POC
// @version 2.0
// @description Sample server to get presigned URL from AWS S3
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	shutdown, _ := otel.InitProvider()
	defer shutdown(ctx)

	router := gin.Default()
	router.Use(otelgin.Middleware("s3-poc"))

	docs.SwaggerInfo.BasePath = "/api"

	group := router.Group("/api")

	group.GET("/s3/files/:name", endpoints.GetFileURL())
	group.GET("/s3/files/otel/:name", endpoints.GetFileURLOtel())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
