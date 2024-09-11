package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"superviseMe/core/entity"
	"superviseMe/core/module"
)

type goalsHandler struct {
	goalsUseCase module.GoalsUseCase
}

func NewGoalsHandler(goalsUseCase module.GoalsUseCase) *goalsHandler {
	return &goalsHandler{goalsUseCase: goalsUseCase}
}

func (e *goalsHandler) CreateGoals(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	goals := &entity.Goals{}
	userGmail := request.Context().Value("email").(string)
	fmt.Println("ini dia:", userGmail)

	goals.PersonalGmail = userGmail
	fmt.Println("ini:", goals.PersonalGmail)

	err := json.NewDecoder(request.Body).Decode(&goals)
	fmt.Println("ini:", goals.PersonalGmail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	goal, err := e.goalsUseCase.CreateGoals(goals)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	responGoal := entity.CreateResponGoal{
		ID:              goal.ID,
		GoalName:        goal.GoalName,
		Supervisor:      goal.SupervisorGmail,
		BackgroundColor: goal.BackgroundColor,
	}

	responsesSucces := entity.ResponsesSucces{Message: "Success", Data: responGoal}
	result, err := json.Marshal(responsesSucces)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	writer.Write(result)
}

func (h *goalsHandler) RequestSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goals := &entity.Goals{}
	userGmail := r.Context().Value("email").(string)
	fmt.Println("ini dia:", userGmail)

	goals.PersonalGmail = userGmail
	fmt.Println("ini:", goals.PersonalGmail)

	supervisorEmail := r.URL.Query().Get("supervisor")
	if supervisorEmail == "" {
		http.Error(w, "Supervisor email is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&goals); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.goalsUseCase.RequestSupervisor(goals.PersonalGmail, supervisorEmail)
	if err != nil {
		http.Error(w, "Failed to request supervisor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Supervisor request sent successfully"))
}

func (h *goalsHandler) AcceptSupervisorRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goals := &entity.Goals{}
	userGmail := r.Context().Value("email").(string)
	fmt.Println("ini dia:", userGmail)

	goals.PersonalGmail = userGmail
	fmt.Println("ini:", goals.PersonalGmail)

	supervisorEmail := r.URL.Query().Get("supervisorEmail")
	if supervisorEmail == "" {
		http.Error(w, "Supervisor email is required", http.StatusBadRequest)
		return
	}

	err := h.goalsUseCase.AcceptSupervisorRequest(userGmail, supervisorEmail)
	if err != nil {
		http.Error(w, "Failed to accept supervisor request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Supervisor request accepted"))
}

func (h *goalsHandler) RejectSupervisorRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goals := &entity.Goals{}
	userGmail := r.Context().Value("email").(string)
	fmt.Println("ini dia:", userGmail)

	goals.PersonalGmail = userGmail
	fmt.Println("ini:", goals.PersonalGmail)

	supervisorEmail := r.URL.Query().Get("supervisorEmail")
	if supervisorEmail == "" {
		http.Error(w, "Supervisor email is required", http.StatusBadRequest)
		return
	}

	err := h.goalsUseCase.RejectSupervisorRequest(userGmail, supervisorEmail)
	if err != nil {
		http.Error(w, "Failed to reject supervisor request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Supervisor request rejected"))
}

// request body create goal
// {
//     "goalName":"mtk",
//     "description":"coba coba",
//     "supervisorGmail":"Samsul@gmail.com",
//     "backgroundColor":"#ffff"
// }

// func (e *goalsHandler) GetGoals(writer http.ResponseWriter, request *http.Request) {
// writer.Header().Set("Content-Type", "application/json")
// goals, err := e.goalsUseCase.GetGoals()
// if err != nil {
// 	writer.Write([]byte(err.Error()))
// 	return
// }
// // responUser := entity.UserResponse{
// // 	ID:       goals.User.ID,
// // 	Name:     goals.User.Name,
// // 	Gmail:    goals.User.Gmail,
// // 	Password: goals.User.Password,
// // }

// var goalsDetailResponses []entity.GoalsDetailRespons
// for _, t := range goals.GoalsDetail {
// 	goalsDetailResponses = append(goalsDetailResponses, entity.GoalsDetailRespons{
// 		ID:          t.ID,
// 		Supervisee:  t.Supervisee,
// 		Supervisor:  t.Supervisor,
// 		GoalsID:     t.GoalsID,
// 		Status:      t.Status,
// 		RequestedAt: t.RequestedAt,
// 		AcceptedAt:  t.AcceptedAt,
// 		RejectedAt:  t.RejectedAt,
// 	})
// }
// responGoals := entity.GoalsResponses{
// 	ID:              goals.ID,
// 	NameGoals:       goals.NameGoals,
// 	Description:     goals.Description,
// 	BackgroundColor: goals.BackgroundColor,
// 	// User:            responUser,
// 	GoalsDetail:     goalsDetailResponses,
// }

// responses := entity.ResponsesSucces{Message: "Succes", Data: responGoals}
// result, err := json.Marshal(responses)
// if err != nil {
// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
// 	errResponse := entity.ResponsesError{Error: err.Error()}
// 	_ = json.NewEncoder(writer).Encode(errResponse)
// 	return
// }
// fmt.Println(responGoals)
// writer.WriteHeader(http.StatusOK)
// writer.Write(result)

// }

// func (e *goalsHandler) GetGoalsByID(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// vars := mux.Vars(request)
// id := vars["id"]
// goals, err := e.goalsUseCase.GetGoalsByID(id)
// if err != nil {
// 	writer.Write([]byte(err.Error()))
// 	return
// }
// responUser := entity.UserResponse{
// 	ID:       goals.User.ID,
// 	Name:     goals.User.Name,
// 	Gmail:    goals.User.Gmail,
// 	Password: goals.User.Password,
// }

// var goalsDetailResponses []entity.GoalsDetailRespons
// for _, t := range goals.GoalsDetail {
// 	goalsDetailResponses = append(goalsDetailResponses, entity.GoalsDetailRespons{
// 		ID:          t.ID,
// 		Supervisee:  t.Supervisee,
// 		Supervisor:  t.Supervisor,
// 		GoalsID:     t.GoalsID,
// 		Status:      t.Status,
// 		RequestedAt: t.RequestedAt,
// 		AcceptedAt:  t.AcceptedAt,
// 		RejectedAt:  t.RejectedAt,
// 	})
// }
// responGoals := entity.GoalsResponses{
// 	ID:              goals.ID,
// 	NameGoals:       goals.NameGoals,
// 	Description:     goals.Description,
// 	BackgroundColor: goals.BackgroundColor,
// 	// User:            responUser,
// 	GoalsDetail: goalsDetailResponses,
// }

// responses := entity.ResponsesSucces{Message: "Succes", Data: responGoals}
// result, err := json.Marshal(responses)
// if err != nil {
// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
// 	return
// }
// writer.WriteHeader(http.StatusOK)
// writer.Write(result)
// }

// func (e *goalsHandler) UpdateGoals(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(request)
// 	id := vars["id"]
// 	var goals *entity.Goals
// 	err := e.goalsUseCase.UpdateGoals(id, goals)
// 	if err != nil {
// 		writer.Write([]byte(err.Error()))
// 		return
// 	}

// 	responsesSucces := entity.ResponsesSucces{Message: "Succes"}
// 	respon, err := json.Marshal(responsesSucces)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		errResponse := entity.ResponsesError{Error: err.Error()}
// 		_ = json.NewEncoder(writer).Encode(errResponse)
// 	}
// 	writer.Write(respon)
// }

// func (e *goalsHandler) DeleteGoals(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(request)
// 	id := vars["id"]
// 	goalsIsActive := request.URL.Query().Get("isActive")
// 	_, err := e.goalsUseCase.DeleteGoals(id, goalsIsActive)
// 	if err != nil {
// 		writer.Write([]byte(err.Error()))
// 		return
// 	}
// 	responsesSucces := entity.Respons{Message: "Success"}
// 	respon, err := json.Marshal(responsesSucces)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		errResponse := entity.ResponsesError{Error: err.Error()}
// 		_ = json.NewEncoder(writer).Encode(errResponse)
// 	}
// 	writer.Write(respon)

// }
