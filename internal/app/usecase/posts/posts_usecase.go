package posts

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"patreon/internal/app"
	"patreon/internal/app/models"
	repoAttaches "patreon/internal/app/repository/attaches"
	repoPosts "patreon/internal/app/repository/posts"
	"patreon/internal/microservices/files/delivery/grpc/client"
	repoFiles "patreon/internal/microservices/files/files/repository/files"
	push_client "patreon/internal/microservices/push/delivery/client"
	"patreon/pkg/utils"

	"github.com/pkg/errors"
)

type PostsUsecase struct {
	repository      repoPosts.Repository
	repositoryData  repoAttaches.Repository
	filesRepository client.FileServiceClient
	imageConvector  utils.ImageConverter
	pusher          push_client.Pusher
}

func NewPostsUsecase(repository repoPosts.Repository, repositoryData repoAttaches.Repository,
	fileClient client.FileServiceClient, pusher push_client.Pusher, convector ...utils.ImageConverter) *PostsUsecase {
	conv := utils.ImageConverter(&utils.ConverterToWebp{})
	if len(convector) != 0 {
		conv = convector[0]
	}
	return &PostsUsecase{
		repository:      repository,
		repositoryData:  repositoryData,
		imageConvector:  conv,
		filesRepository: fileClient,
		pusher:          pusher,
	}
}

// GetPosts Errors:
// 		app.GeneralError with Errors:
// 			repository.DefaultErrDB
func (usecase *PostsUsecase) GetPosts(creatorId int64, userId int64,
	pag *models.Pagination, withDraft bool) ([]models.Post, error) {
	return usecase.repository.GetPosts(creatorId, userId, pag, withDraft)
}

// GetAvailablePosts Errors:
// 		app.GeneralError with Errors:
// 			repository.DefaultErrDB
func (usecase *PostsUsecase) GetAvailablePosts(userID int64, pag *models.Pagination) ([]models.AvailablePost, error) {
	return usecase.repository.GetAvailablePosts(userID, pag)
}

// GetPost Errors:
//		repository.NotFound
// 		app.GeneralError with Errors:
// 			repository.DefaultErrDB
func (usecase *PostsUsecase) GetPost(postId int64, userId int64, addView bool) (*models.PostWithAttach, error) {
	post, err := usecase.repository.GetPost(postId, userId, addView)
	if err != nil {
		return nil, err
	}
	res := &models.PostWithAttach{Post: post, Data: []models.AttachWithoutLevel{}}
	res.Data, err = usecase.repositoryData.GetAttaches(postId)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Delete Errors:
// 		app.GeneralError with Errors:
// 			repository.DefaultErrDB
func (usecase *PostsUsecase) Delete(postId int64) error {
	return usecase.repository.Delete(postId)
}

// Update Errors:
// 		repository.NotFound
//		models.InvalidAwardsId
//		models.EmptyTitle
//		app.GeneralError with Errors:
//			app.UnknownError
//			repository.DefaultErrDB
func (usecase *PostsUsecase) Update(log *logrus.Entry, post *models.UpdatePost) error {
	if err := post.Validate(); err != nil {
		if errors.Is(err, models.EmptyTitle) || errors.Is(err, models.InvalidAwardsId) {
			if post.IsDraft && errors.Is(err, models.EmptyTitle) {
				return usecase.repository.UpdatePost(post)
			}
			return err
		}
		return &app.GeneralError{
			Err:         app.UnknownError,
			ExternalErr: errors.Wrap(err, "failed process of validation creator"),
		}
	}

	if !post.IsDraft {
		if oldPost, err := usecase.repository.GetPost(post.ID, EmptyUser, false); err == nil {
			if oldPost.IsDraft {
				if creatorId, err := usecase.repository.GetPostCreator(post.ID); err == nil {
					errPush := usecase.pusher.NewPost(creatorId, post.ID, post.Title)
					if errPush != nil {
						log.Errorf("Try push new post, and got err %s", errPush)
					}
				} else {
					log.Errorf("Try get cretor post, and got err %s", err)
				}
			}
		} else {
			log.Errorf("Try get cretor old post, and got err %s", err)
		}
	}

	return usecase.repository.UpdatePost(post)
}

// Create Errors:
//		models.InvalidAwardsId
//		models.InvalidCreatorId
//		models.EmptyTitle
//		app.GeneralError with Errors:
//			app.UnknownError
//			repository.DefaultErrDB
func (usecase *PostsUsecase) Create(log *logrus.Entry, post *models.CreatePost) (int64, error) {
	if err := post.Validate(); err != nil {
		if errors.Is(err, models.EmptyTitle) || errors.Is(err, models.InvalidCreatorId) ||
			errors.Is(err, models.InvalidAwardsId) {
			if errors.Is(err, models.EmptyTitle) && post.IsDraft {
				return usecase.repository.Create(post)
			}
			return app.InvalidInt, err
		}
		return app.InvalidInt, &app.GeneralError{
			Err:         app.UnknownError,
			ExternalErr: errors.Wrap(err, "failed process of validation creator"),
		}
	}
	postId, err := usecase.repository.Create(post)
	if !post.IsDraft {
		errPush := usecase.pusher.NewPost(post.CreatorId, postId, post.Title)
		if errPush != nil {
			log.Errorf("Try push new post, and got err %s", errPush)
		}
	}
	return postId, err
}

// GetCreatorId Errors:
//  	repository.NotFound
//  	app.GeneralError with Errors:
//   		repository.DefaultErrDB
func (usecase *PostsUsecase) GetCreatorId(postId int64) (int64, error) {
	aw, err := usecase.repository.GetPostCreator(postId)
	if err != nil {
		return app.InvalidInt, err
	}
	return aw, nil
}

// LoadCover Errors:
//		repository.NotFound
//		app.GeneralError with Errors:
//			repository.DefaultErrDB
//			repository_os.ErrorCreate
//   		repository_os.ErrorCopyFile
//			utils.ConvertErr
//  		utils.UnknownExtOfFileName
func (usecase *PostsUsecase) LoadCover(data io.Reader, name repoFiles.FileName, postId int64) error {
	if _, err := usecase.repository.GetPostCreator(postId); err != nil {
		return err
	}

	var err error
	data, name, err = usecase.imageConvector.Convert(context.Background(), data, name)
	if err != nil {
		return errors.Wrap(err, "failed convert to webp of update post cover")
	}

	path, err := usecase.filesRepository.SaveFile(context.Background(), data, name, repoFiles.Image)
	if err != nil {
		return err
	}

	return usecase.repository.UpdateCoverPost(postId, app.LoadFileUrl+path)
}
