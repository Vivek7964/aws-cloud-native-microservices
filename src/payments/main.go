package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Vivek7964/microservice-demo/payments/config"
	"github.com/Vivek7964/microservice-demo/payments/controller"
	_ "github.com/Vivek7964/microservice-demo/payments/docs"
	"github.com/sethvargo/go-envconfig/pkg/envconfig"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Payments API
// @version 1.0
// @description This API serves the product payments

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	ctx := context.Background()

	var config config.AppConfiguration
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	c, err := controller.NewController(config)
	if err != nil {
		log.Fatalln("Error creating controller", err)
	}

	payments := r.Group("/payments")
	{
		payments.GET("/client-secret/:id", c.GetPaymentIntentByCartID)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.Run(":" + strconv.Itoa(config.Port))
}