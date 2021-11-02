package usecase_factory

import (
	repCsrf "patreon/internal/app/csrf/repository/jwt"
	repAccess "patreon/internal/app/repository/access"
	repoAwrds "patreon/internal/app/repository/awards"
	repCreator "patreon/internal/app/repository/creator"
	repoFiles "patreon/internal/app/repository/files"
	repoLikes "patreon/internal/app/repository/likes"
	repoPosts "patreon/internal/app/repository/posts"
	repoPostsData "patreon/internal/app/repository/posts_data"
	useSubscr "patreon/internal/app/repository/subscribers"
	repUser "patreon/internal/app/repository/user"

	"patreon/internal/app/sessions"
)

//go:generate mockgen -destination=mocks/mock_repository_factory.go -package=mock_repository_factory . RepositoryFactory

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetCreatorRepository() repCreator.Repository
	GetAwardsRepository() repoAwrds.Repository
	GetCsrfRepository() repCsrf.Repository
	GetSessionRepository() sessions.SessionRepository
	GetAccessRepository() repAccess.Repository
	GetSubscribersRepository() useSubscr.Repository
	GetPostsRepository() repoPosts.Repository
	GetLikesRepository() repoLikes.Repository
	GetPostsDataRepository() repoPostsData.Repository
	GetFilesRepository() repoFiles.Repository
}
