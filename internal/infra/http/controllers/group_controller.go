package controllers

import (
	"errors"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
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
// @Description Get the average transparency detected by a particular sensors in a group.
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
