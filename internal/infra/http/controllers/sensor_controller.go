package controllers

import (
	"errors"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type SensorController struct {
	sensorService app.SensorService
}

func NewSensorController(s app.SensorService) SensorController {
	return SensorController{
		sensorService: s,
	}
}

// @Summary Get Sensor Temperature Average
// @Description Get the average temperature detected by a particular sensor between specified date/time pairs.
// @ID get-sensor-temperature-average
// @Param codeName path string true "Code Name of the Sensor"
// @Param from query int true "From Date/Time (UNIX Timestamp)"
// @Param till query int true "Till Date/Time (UNIX Timestamp)"
// @Produce json
// @Success 200 {object} int64
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /sensor/{codeName}/temperature/average [get]
func (c SensorController) GetSensorTemperatureAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		codeName := chi.URLParam(r, "codeName")
		if codeName == "" {
			err := errors.New("code name is empty")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		fromStr := r.URL.Query().Get("from")
		tillStr := r.URL.Query().Get("till")

		if fromStr == "" || tillStr == "" {
			err := errors.New("from or till parameter is missing")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		from, err := strconv.ParseInt(fromStr, 10, 64)
		if err != nil {
			err = errors.New("from parameter is not a valid UNIX timestamp")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		till, err := strconv.ParseInt(tillStr, 10, 64)
		if err != nil {
			err = errors.New("till parameter is not a valid UNIX timestamp")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		if from < 0 || till < 0 || from > till {
			err = errors.New("from and till parameters are not in a valid range")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		//addr, err = c.addrService.Update(addr)

		//
		//var addrDto resources.AddressDto
		averageTemperature, err := c.sensorService.GetAverageTemperatureBySensor(codeName, from, till)
		if err != nil {
			log.Printf("SensorController: %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, averageTemperature)

	}
}
