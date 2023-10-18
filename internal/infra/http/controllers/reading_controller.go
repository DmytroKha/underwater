package controllers

import "github.com/DmytroKha/underwater/internal/app"

type ReadingController struct {
	readingService app.ReadingService
	sensorService  app.SensorService
}

func NewReadingController(s app.ReadingService, ss app.SensorService) ReadingController {
	return ReadingController{
		readingService: s,
		sensorService:  ss,
	}
}
