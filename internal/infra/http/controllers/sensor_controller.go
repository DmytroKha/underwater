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

		averageTemperature, err := c.sensorService.GetAverageTemperatureBySensor(codeName, from, till)
		if err != nil {
			log.Printf("SensorController: %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, averageTemperature)

	}
}

// @Summary Get Region Min Temperature
// @Description Get the minimum temperature detected by sensors in the region.
// @ID get-region-min-temperature
// @Param xMin query int true "Min X"
// @Param xMax query int true "Max X"
// @Param yMin query int true "Min Y"
// @Param yMax query int true "Max Y"
// @Param zMin query int true "Min Z"
// @Param zMax query int true "Max Z"
// @Produce json
// @Success 200 {object} int64
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /region/temperature/min [get]
func (c SensorController) GetRegionMinTemperature() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		xMinStr := r.URL.Query().Get("xMin")
		xMaxStr := r.URL.Query().Get("xMax")

		yMinStr := r.URL.Query().Get("yMin")
		yMaxStr := r.URL.Query().Get("yMax")

		zMinStr := r.URL.Query().Get("zMin")
		zMaxStr := r.URL.Query().Get("zMax")

		if xMinStr == "" || xMaxStr == "" || yMinStr == "" || yMaxStr == "" || zMinStr == "" || zMaxStr == "" {
			err := errors.New("some region parameters are missing")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		xMin, err := strconv.ParseFloat(xMinStr, 64)
		if err != nil {
			err = errors.New("xMin parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		xMax, err := strconv.ParseFloat(xMaxStr, 64)
		if err != nil {
			err = errors.New("xMax parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		yMin, err := strconv.ParseFloat(yMinStr, 64)
		if err != nil {
			err = errors.New("yMin parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		yMax, err := strconv.ParseFloat(yMaxStr, 64)
		if err != nil {
			err = errors.New("yMax parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		zMin, err := strconv.ParseFloat(zMinStr, 64)
		if err != nil {
			err = errors.New("zMin parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		zMax, err := strconv.ParseFloat(zMaxStr, 64)
		if err != nil {
			err = errors.New("zMax parameter is not valid")
			log.Printf("SensorController: %s", err)
			BadRequest(w, err)
			return
		}

		regionMinTemperature, err := c.sensorService.GetRegionMinTemperature(xMin, yMin, zMin, xMax, yMax, zMax)
		if err != nil {
			log.Printf("SensorController: %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, regionMinTemperature)

	}
}
