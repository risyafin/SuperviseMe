package commentrepository

import (
	"fmt"
	"superviseMe/core/entity"
	commentRepo "superviseMe/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) commentRepo.CommentRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateComment(comment *entity.Comment) error {
	err := r.DB.Create(&comment).Error
	if err != nil {
		// Log the error (opsional)
		fmt.Println("Failed to create comment:", err)
		return err
	}
	return nil
}
