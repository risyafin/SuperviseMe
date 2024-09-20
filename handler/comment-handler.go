package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"superviseMe/core/entity"
	"superviseMe/core/module"

	"github.com/gorilla/mux"
)

type commentHandler struct {
	commentUsecase module.CommentUsecase
}

func NewCommentHandler(commentUsecase module.CommentUsecase) *commentHandler {
	return &commentHandler{commentUsecase: commentUsecase}
}

func (h *commentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	// Ambil card_id dari URL parameter
	vars := mux.Vars(r)
	cardID, err := strconv.Atoi(vars["card_id"])
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return
	}

	// Ambil user_id dari context (misalnya melalui middleware auth)
	userID := r.Context().Value("userID").(int)
	fmt.Println("ini dia:", userID)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parsing request body
	var req entity.Comment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Panggil usecase untuk membuat comment
	err = h.commentUsecase.CreateComment(cardID, userID, req.Message)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment created successfully"})
}
