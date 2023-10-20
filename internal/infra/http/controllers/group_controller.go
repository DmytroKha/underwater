package controllers

import (
	"errors"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/DmytroKha/underwater/internal/infra/http/resources"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type GroupController struct {
	groupService app.GroupService
}

func NewGroupController(s app.GroupService) GroupController {
	return GroupController{
		groupService: s,
	}
}

// @Summary Get Group Temperature Average
// @Description Get the average temperature detected by a particular sensors in a group.
// @ID get-group-temperature-average
// @Param groupName path string true "Group Name of the Sensors"
// @Produce json
// @Success 200 {object} int64
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /group/{groupName}/temperature/average [get]
func (c GroupController) GetGroupTemperatureAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := chi.URLParam(r, "groupName")
		if groupName == "" {
			err := errors.New("group name is empty")
			log.Printf("GroupController: %s", err)
			BadRequest(w, err)
			return
		}

		averageTemperature, err := c.groupService.GetAverageTemperatureForGroup(groupName)
		if err != nil {
			log.Printf("GroupController: %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, averageTemperature)

	}
}

// @Summary Get Group Transparency Average
// @Description Get the average transparency detected by sensors in a group.
// @ID get-group-transparency-average
// @Param groupName path string true "Group Name of the Sensors"
// @Produce json
// @Success 200 {object} int64
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /group/{groupName}/transparency/average [get]
func (c GroupController) GetGroupTransparencyAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := chi.URLParam(r, "groupName")
		if groupName == "" {
			err := errors.New("group name is empty")
			log.Printf("GroupController: %s", err)
			BadRequest(w, err)
			return
		}

		averageTemperature, err := c.groupService.GetAverageTransparencyForGroup(groupName)
		if err != nil {
			log.Printf("GroupController: %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, averageTemperature)

	}
}

// @Summary Get Group Fish Species
// @Description Get fish species in a group.
// @ID get-group-fish-species
// @Param groupName path string true "Group Name of the Sensors"
// @Produce json
// @Success 200 {object} []resources.FishDto
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /group/{groupName}/species [get]
func (c GroupController) GetGroupFishSpecies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := chi.URLParam(r, "groupName")
		if groupName == "" {
			err := errors.New("group name is empty")
			log.Printf("GroupController: %s", err)
			BadRequest(w, err)
			return
		}

		fishes, err := c.groupService.GetFishesForGroup(groupName)
		if err != nil {
			log.Printf("GroupController: %s", err)
			InternalServerError(w, err)
			return
		}

		var fishesDto resources.FishDto
		Success(w, fishesDto.DomainToDtoCollection(fishes))

	}
}

// @Summary Get Group Top Fish Species
// @Description Get top fish species in a group.
// @ID get-group-top-fish-species
// @Param groupName path string true "Group Name of the Sensors"
// @Param N path string true "Count of top fishes"
// @Param from query int false "From Date/Time (UNIX Timestamp)"
// @Param till query int false "Till Date/Time (UNIX Timestamp)"
// @Produce json
// @Success 200 {object} []resources.FishDto
// @Failure 400 {string} http.StatusBadRequest
// @Failure 500 {string} http.StatusInternalServerError
// @Router /group/{groupName}/species/top/{N} [get]
func (c GroupController) GetGroupTopFishSpecies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupName := chi.URLParam(r, "groupName")
		if groupName == "" {
			err := errors.New("group name is empty")
			log.Printf("GroupController: %s", err)
			BadRequest(w, err)
			return
		}

		fishCount, err := strconv.ParseInt(chi.URLParam(r, "N"), 10, 64)
		if err != nil {
			log.Printf("GroupController: %s", err)
			BadRequest(w, err)
			return
		}

		var from, till int64

		fromStr := r.URL.Query().Get("from")
		tillStr := r.URL.Query().Get("till")

		if fromStr != "" || tillStr != "" {
			from, err = strconv.ParseInt(fromStr, 10, 64)
			if err != nil {
				err = errors.New("from parameter is not a valid UNIX timestamp")
				log.Printf("GroupController: %s", err)
				BadRequest(w, err)
				return
			}

			till, err = strconv.ParseInt(tillStr, 10, 64)
			if err != nil {
				err = errors.New("till parameter is not a valid UNIX timestamp")
				log.Printf("GroupController: %s", err)
				BadRequest(w, err)
				return
			}

			if from < 0 || till < 0 || from > till {
				err = errors.New("from and till parameters are not in a valid range")
				log.Printf("GroupController: %s", err)
				BadRequest(w, err)
				return
			}
		}

		fishes, err := c.groupService.GetTopFishesForGroup(groupName, fishCount, from, till)
		if err != nil {
			log.Printf("GroupController: %s", err)
			InternalServerError(w, err)
			return
		}

		var fishesDto resources.FishDto
		Success(w, fishesDto.DomainToDtoCollection(fishes))

	}
}
