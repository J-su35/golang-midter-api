package main

import (
	"context"
	"disbursement-api/internal/auth"
	"disbursement-api/internal/item"
	"disbursement-api/internal/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pwd := os.Getenv("PWD")
	port := os.Getenv("PORT")


	db, err := gorm.Open(
		postgres.Open(
			"postgres://postgres:"+pwd+"@localhost:5432/task",
		),
	)
	if err != nil {
		log.Panic(err)
	}
	controller := item.NewController(db)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}
	r.Use(cors.New(config))

	userController := user.NewController(db, os.Getenv("JWT_SECRET"))
	r.POST("/login", userController.Login)

	//Register router
	items := r.Group("/items")
	items.Use(auth.Guard(os.Getenv("JWT_SECRET")))
	{
		items.POST("", controller.CreateItem)
		items.GET("", controller.FindItems)
		items.PATCH("/:id", controller.UpdateItemStatus)
		items.GET("/:id", controller.FindItemById)
		items.DELETE("/:id", controller.DeleteItemById)
		items.PUT("/:id", controller.ReplaceItem)
	}
	

	//Graceful shutdown regular method 
	srv := &http.Server{
        Addr:    ":"+port,
        Handler: r.Handler(),
    }
    go func() {
        // service connections
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()
    // Wait for interrupt signal to gracefully shutdown the server with
    // a timeout of ___ seconds.
    quit := make(chan os.Signal, 1)
    // kill (no param) default send syscall.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown Server ...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
    // catching ctx.Done(). timeout of 5 seconds.
    <-ctx.Done()
	log.Println("timeout of 5 seconds.")
    log.Println("Server exiting")
	//Graceful shutdown regular method end
}