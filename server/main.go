package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/initialize"
	router "github.com/threeforone/order-meal/server/router"
)

func init() {
	global.DB = initialize.Mysql()
}

func main() {
	//close DB connecting
	defer global.CloseDB()

	g := gin.Default()
	router.Routers(g)

	// create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Conf.App.Port),
		Handler: g,
	}

	//start up HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//wait INT signal or TERM signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//stop HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
