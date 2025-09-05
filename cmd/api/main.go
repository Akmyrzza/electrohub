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

	"github.com/akmyrzza/electrohub/internal/products/delivery"
	"github.com/akmyrzza/electrohub/internal/products/repository"
	"github.com/akmyrzza/electrohub/internal/products/usecase"
	"github.com/go-chi/chi/v5"
)

var buildVersion = "dev"

func main() {
	fmt.Println("Electrohub API starting...")

	productRepository := repository.NewInMemoryProductRepository()
	productUseCase := usecase.NewProductService(productRepository)
	productHandler := delivery.NewProductHandler(productUseCase)

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Route("/products", func(pr chi.Router) {
			pr.Get("/", productHandler.ListProducts)
			pr.Post("/", productHandler.CreateProduct)
		})
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Electrohub API (version: %s) running on %s\n", buildVersion, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
