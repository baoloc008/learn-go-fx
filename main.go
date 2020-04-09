package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

type ServerConfig struct {
	listenAddr string
}

func NewServerConfig() (*ServerConfig, error) {
	//return &ServerConfig{listenAddr: ":9620"}, errors.New("load config failed")
	return &ServerConfig{listenAddr: ":9620"}, nil
}

func main() {
	app := fx.New(
		fx.Provide(NewServerConfig),
		fx.Invoke(startServer),
	)
	app.Run()
}

func startServer(lifecycle fx.Lifecycle, config *ServerConfig) {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	srv := http.Server{
		Addr:    config.listenAddr,
		Handler: router,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatalf("Server ListenAndServe Error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server is shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Fatalf("Server Shutdown Error: %v", err)
			} else {
				log.Println("Server Shutdown Properly")
			}
			return err
		},
	})
}
