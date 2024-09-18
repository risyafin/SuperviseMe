package handler

import (
	"encoding/json"
	"net/http"
	"superviseMe/core/entity"
	"superviseMe/core/module"
)

type notificationHandler struct {
	notificationUseCase module.NotificationUsecase
}

func NewNotificationHandler(notificationUseCase module.NotificationUsecase) *notificationHandler {
	return &notificationHandler{notificationUseCase: notificationUseCase}
}
func (e *notificationHandler) GetNotification(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userGmail, ok := request.Context().Value("email").(string)
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		errResponse := entity.ResponsesError{Error: "Invalid gmail in context"}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}

	notification, err := e.notificationUseCase.GetNotification(userGmail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}

	var typeNotificationRespon []entity.TypeNotificationRespon
	for _, s := range notification.TypeNotification {
		typeNotificationRespon = append(typeNotificationRespon, entity.TypeNotificationRespon{
			ID:             s.ID,
			Name:           s.Name,
			NotificationID: s.NotificationID,
		})
	}

	responNotification := entity.NotificationRespon{
		ID:               notification.ID,
		PersonalEmail:    notification.PersonalEmail,
		SupervisorEmail:  notification.SupervisorEmail,
		GoalsID:          notification.GoalsID,
		Message:          notification.Message,
		Status:           notification.Status,
		TypeNotification: typeNotificationRespon,
	}

	responses := entity.ResponsesSucces{Message: "Succes", Data: responNotification}
	result, err := json.Marshal(responses)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(result)
}
