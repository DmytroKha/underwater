package http

import (
	"github.com/DmytroKha/underwater/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Router(
	sensorController controllers.SensorController,
	groupController controllers.GroupController,
) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api", func(apiRouter chi.Router) {

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			// Health
			apiRouter.Route("/ping", func(healthRouter chi.Router) {
				healthRouter.Get("/", PingHandler())

				healthRouter.Handle("/*", NotFoundJSON())
			})

			// Routes
			apiRouter.Group(func(apiRouter chi.Router) {

				SensorRouter(apiRouter, sensorController)
				GroupRouter(apiRouter, groupController)

				apiRouter.Handle("/*", NotFoundJSON())
			})

		})

		// Serve Swagger UI at /swagger
		apiRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("doc.json"), // Point to your Swagger JSON URL
		))
	})

	return router
}

func SensorRouter(r chi.Router, sensorController controllers.SensorController) {
	r.Route("/sensor", func(sensorRouter chi.Router) {
		sensorRouter.Get("/{codeName}/temperature/average", sensorController.GetSensorTemperatureAverage())
	})
}

func GroupRouter(r chi.Router, groupController controllers.GroupController) {
	r.Route("/group", func(sensorRouter chi.Router) {
		sensorRouter.Get("/{groupName}/temperature/average", groupController.GetGroupTemperatureAverage())
		sensorRouter.Get("/{groupName}/transparency/average", groupController.GetGroupTransparencyAverage())
		sensorRouter.Get("/{groupName}/species", groupController.GetGroupFishSpecies())
		//sensorRouter.Get("/{groupName}/species/top/{N}", groupController.GetGroupTopFishSpecies())
	})
}
