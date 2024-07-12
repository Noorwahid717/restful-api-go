package dto

import (
	"time"

	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CreateEvent struct {
	Name        string    `json:"name" binding:"required,alpha,min=3,max=20"`
	Description string    `json:"description" binding:"required,alpha,min=3,max=20"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required"`
	Location    string    `json:"location" binding:"required,alpha,min=3,max=20"`
}

type UpdateEvent struct {
	Name        string    `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	Description string    `json:"description,omitempty" binding:"alpha,min=3,max=20"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	Location    string    `json:"location,omitempty" binding:"alpha,min=3,max=20"`
}

type EventResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
}

func ToEventResponse(from dto.Event) EventResponse {
	return EventResponse{
		Id:          from.Id,
		Name:        from.Name,
		Description: from.Description,
		StartDate:   from.StartDate,
		EndDate:     from.EndDate,
		Location:    from.Location,
	}
}

// func ToEventListResponse(from []dto.Event) []EventResponse {
// 	var result []EventResponse
// 	for _, item := range from {
// 		result = append(result, ToEventResponse(item))
// 	}
// 	return result
// }

func ToCreateEvent(from CreateEvent) dto.CreateEvent {
	return dto.CreateEvent{
		Name:        from.Name,
		Description: from.Description,
		StartDate:   from.StartDate,
		EndDate:     from.EndDate,
		Location:    from.Location,
	}
}

func ToUpdateEvent(from UpdateEvent) dto.UpdateEvent {
	return dto.UpdateEvent{
		Name:        from.Name,
		Description: from.Description,
		StartDate:   from.StartDate,
		EndDate:     from.EndDate,
		Location:    from.Location,
	}
}
