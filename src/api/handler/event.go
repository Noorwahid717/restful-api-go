package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/dependency"
	_ "github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

type EventHandler struct {
	usecase *usecase.EventUsecase
}

func NewEventHandler(cfg *config.Config) *EventHandler {
	return &EventHandler{
		usecase: usecase.NewEventUsecase(cfg, dependency.GetEventRepository(cfg)),
	}
}

// CreateEvent godoc
// @Summary Create a Event
// @Description Create a Event
// @Tags Events
// @Accept json
// @produces json
// @Param Request body dto.CreateEvent true "Create a Event"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.EventResponse} "Event response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/events/ [post]
// @Security AuthBearer
func (h *EventHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateEvent, dto.ToEventResponse, h.usecase.Create)
}

// UpdateEvent godoc
// @Summary Update a Event
// @Description Update a Event
// @Tags Events
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateEvent true "Update a Event"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.EventResponse} "Event response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/events/{id} [put]
// @Security AuthBearer
func (h *EventHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateEvent, dto.ToEventResponse, h.usecase.Update)
}

// DeleteEvent godoc
// @Summary Delete a Event
// @Description Delete a Event
// @Tags Events
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 204 {object} helper.BaseHttpResponse "No content"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/events/{id} [delete]
// @Security AuthBearer
func (h *EventHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetEvent godoc
// @Summary Get a Event
// @Description Get a Event
// @Tags Events
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.EventResponse} "Event response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/events/{id} [get]
func (h *EventHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToEventResponse, h.usecase.GetById)
}

// GetEvents godoc
// @Summary Get Events
// @Description Get Events
// @Tags Events
// @Accept json
// @produces json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param filter query string false "Filter"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.EventResponse]} "Event response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/events/ [get]
