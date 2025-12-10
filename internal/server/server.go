package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mrxacker/go-myapp/internal/config"
	"github.com/mrxacker/go-myapp/internal/handlers"
	"github.com/mrxacker/go-myapp/internal/service"
	"github.com/mrxacker/go-myapp/pkg/logger"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config, userService *service.UserService) *Server {
	r := chi.NewRouter()

	// Chi middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// Custom middleware
	r.Use(handlers.LoggingMiddleware)

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health check
	r.Get("/health", handlers.HealthCheck)

	// API routes
	r.Route("/api", func(r chi.Router) {
		// User routes
		userHandler := handlers.NewUserHandler(userService)
		r.Route("/users", func(r chi.Router) {
			//r.Get("/", userHandler.List)
			r.Post("/", userHandler.Create)
			r.Get("/{id}", userHandler.Get)
			//r.Put("/{id}", userHandler.Update)
			//r.Delete("/{id}", userHandler.Delete)
		})
	})

	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.HTTPPort),
			Handler:      r,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Get().Info("Shutting down HTTP server")
	return s.httpServer.Shutdown(ctx)
}
