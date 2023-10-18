package http

import (
	"github.com/DmytroKha/underwater/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func Router(
	sensorController controllers.SensorController,
	readingController controllers.ReadingController,
) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api", func(apiRouter chi.Router) {

		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			// Routes
			apiRouter.Group(func(apiRouter chi.Router) {

				SensorRouter(apiRouter, sensorController)
				//ReadingRouter(apiRouter, readingController)

				apiRouter.Handle("/*", NotFoundJSON())
			})
		})
	})

	return router
}

func SensorRouter(r chi.Router, sensorController controllers.SensorController) {
	r.Route("/sensor", func(sensorRouter chi.Router) {
		sensorRouter.Get("/{codeName}/temperature/average", sensorController.GetSensorTemperatureAverage())
	})
}
