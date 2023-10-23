package jobs

import (
	"github.com/DmytroKha/underwater/internal/app"
	"log"
)

func StartSensorDataGeneration(readingService *app.ReadingService, sensorService *app.SensorService) {
	sensors, err := (*sensorService).FindAll()
	if err != nil {
		log.Print(err)
		return
	}

	for _, sensor := range sensors {
		go (*readingService).GenerateSensorData(sensor)
	}
	select {}
}
