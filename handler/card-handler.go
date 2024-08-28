package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"superviseMe/core/entity"
	"superviseMe/core/module"
)

type cardHandler struct {
	cardUsecase module.CardUsecase
}

func NewCardHandler(cardUsecase module.CardUsecase) *cardHandler {
	return &cardHandler{cardUsecase: cardUsecase}
}

func (e *cardHandler) CreateGoals(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	card := &entity.Card{}
	userGmail := request.Context().Value("gmail").(string)
	fmt.Println("ini dia:", userGmail)

	card.GmailPersonal = userGmail
	fmt.Println("ini:", card.GmailPersonal)

	err := json.NewDecoder(request.Body).Decode(&card)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	card, err = e.cardUsecase.CreateCard(card)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}
	var checkListCard []entity.CheckListCardResponse
	for _, s := range card.CheckListCard {
		checkListCard = append(checkListCard, entity.CheckListCardResponse{
			ID:     s.ID,
			CardID: s.CardID,
			Name:   s.Name,
			IsDone: s.IsDone,
		})
	}

	responCard := entity.ResponCreateCard{
		ID:            card.ID,
		CardName:      card.CardName,
		Label:         card.Label,
		Attachment:    card.Attachment,
		Content:       card.Content,
		StartDate:     card.StartDate,
		EndDate:       card.EndDate,
		CheckListCard: checkListCard,
		ListID:        card.ListID,
	}

	responsesSucces := entity.ResponsesSucces{Message: "Success", Data: responCard}
	result, err := json.Marshal(responsesSucces)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	writer.Write(result)
}

// reques body
// {
// 	"NameCard": "My First Card",
// 	"checklists": [
// 	  {
// 		"name": "Checklist 1"
// 	  },
// 	  {
// 		"name": "Checklist 2"
// 	  },
// 	  {
// 		"name": "Checklist 3"
// 	  }
// 	]
//   }
