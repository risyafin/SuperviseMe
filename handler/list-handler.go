package handler

import (
	"encoding/json"
	"net/http"
	"superviseMe/core/entity"
	"superviseMe/core/module"
)

type listHandler struct {
	listUseCase module.ListUsecase
}

func NewListHandler(listUseCase module.ListUsecase) *listHandler {
	return &listHandler{listUseCase: listUseCase}
}
func (e *listHandler) GetList(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userGmail, ok := request.Context().Value("email").(string)
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		errResponse := entity.ResponsesError{Error: "Invalid gmail in context"}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}

	list, err := e.listUseCase.GetList(userGmail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}

	cardRespon := make([]entity.CardRespons, len(list.Card))
	for i, card := range list.Card {
		cardRespon[i] = entity.CardRespons{
			ID:        card.ID,
			CardName:  card.CardName,
			Label:     card.Label,
			StartDate: card.StartDate,
			EndDate:   card.EndDate,
			ListID:    card.ListID,
		}
	}

	responList := entity.ListRespon{
		ID:                list.ID,
		ListName:          list.ListName,
		Card:              cardRespon,
	}

	responses := entity.ResponsesSucces{Message: "Succes", Data: responList}
	if err := json.NewEncoder(writer).Encode(responses); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
