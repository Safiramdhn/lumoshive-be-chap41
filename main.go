package main

import (
	"context"
	"log"
	"lumoshive-be-chap41/infra"
	"lumoshive-be-chap41/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxt); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
