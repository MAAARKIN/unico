package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MAAARKIN/unico/api/handler"
	"github.com/MAAARKIN/unico/config"
	"github.com/MAAARKIN/unico/container"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/MAAARKIN/unico/api/docs"
)

var Server *http.Server

type Router interface {
	Route(r chi.Router)
}

type RouterContainer struct {
	Health Router
	Feira  Router
}

func NewRouterContainer(cdi container.Dependency) RouterContainer {

	return RouterContainer{
		Health: handler.Health{},
		Feira:  handler.NewFeiraHandler(cdi.Services.FeiraService),
	}
}

func addCors() func(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposedHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials:   true,
		MaxAge:             300, // Maximum value not ignored by any of major browsers
		OptionsPassthrough: false,
	}).Handler
}

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// @BasePath /v1

// @securityDefinitions.apikey carrierIdAuthentication
// @in header
// @name carrierId
func StartHttpServer(cfg config.Config, cdi container.Dependency) {
	c := chi.NewRouter()

	routes := NewRouterContainer(cdi)
	c.Use(addCors())

	c.Route("/v1", func(r chi.Router) {

		r.Use(chiMiddleware.Logger)
		r.Use(chiMiddleware.Recoverer)
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Route("/health", routes.Health.Route)
		r.Route("/feiras", routes.Feira.Route)
	})

	c.Get("/swagger/*", httpSwagger.WrapHandler)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      c,
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 25 * time.Second,
	}

	Server = s

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}()
	log.Printf("API Up")
}
