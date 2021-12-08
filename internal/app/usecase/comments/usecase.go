package usecase_comments

import (
	"patreon/internal/app/models"
)

//go:generate mockgen -destination=mocks/mock_comments_usecase.go -package=mock_usecase -mock_names=Usecase=CommentsUsecase . Usecase

type Usecase interface {
	// Create Errors:
	//		repository_postgresql.CommentAlreadyExist
	//		models.InvalidPostId
	//		models.InvalidUserId
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	Create(cm *models.Comment) (int64, error)

	// Update Errors:
	//		repository.NotFound
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	Update(cm *models.Comment) error

	// CheckExists Errors:
	//		repository_postgresql.CommentAlreadyExist
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	CheckExists(commentId int64) error

	// GetUserComments Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetUserComments(userId int64, pag *models.Pagination) ([]models.UserComment, error)

	// GetPostComments Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetPostComments(postId int64, pag *models.Pagination) ([]models.PostComment, error)

	// Delete Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	Delete(commentId int64) error
}