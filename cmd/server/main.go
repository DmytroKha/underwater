package main

import (
	"context"
	"fmt"
	"github.com/DmytroKha/underwater/config"
	_ "github.com/DmytroKha/underwater/docs"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/DmytroKha/underwater/internal/infra/database"
	"github.com/DmytroKha/underwater/internal/infra/http"
	"github.com/DmytroKha/underwater/internal/infra/http/controllers"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

// @title       Halo Underwater API
// @version     1.0
// @description API Server for Halo Underwater application.
// @host     localhost:8080
// @BasePath  /api/v1
func main() {

	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	var conf = config.GetConfiguration()

	err := database.Migrate(conf)
	if err != nil {
		log.Fatalf("Unable to apply migrations: %q\n", err)
	}

	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}

	//Repository
	sensorRepository := database.NewSensorRepository(sess)
	readingRepository := database.NewReadingRepository(sess)
	fishSpeciesRepository := database.NewFishSpeciesRepository(sess)

	//Service
	//var fishSpeciesService app.FishSpeciesService
	sensorService := app.NewSensorService(sensorRepository)
	fishSpeciesService := app.NewFishSpeciesService(fishSpeciesRepository)
	readingService := app.NewReadingService(readingRepository, sensorService, fishSpeciesService)

	//Controllers
	sensorController := controllers.NewSensorController(sensorService)
	readingController := controllers.NewReadingController(readingService, sensorService)

	// HTTP Server
	err = http.Server(
		ctx,
		http.Router(
			sensorController,
			readingController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}

	readingService.StartSensorDataGeneration()

}
