package repository_factory

import (
	"patreon/internal/app"
	repCsrf "patreon/internal/app/csrf/repository/jwt"
	repository_access "patreon/internal/app/repository/access"
	repoAwrds "patreon/internal/app/repository/awards"
	repAwardsPsql "patreon/internal/app/repository/awards/postgresql"
	repCreator "patreon/internal/app/repository/creator"
	repCreatorPsql "patreon/internal/app/repository/creator/postgresql"
	repository_subscribers "patreon/internal/app/repository/subscribers"
	repoLikes "patreon/internal/app/repository/likes"
	repLikesPsql "patreon/internal/app/repository/likes/postgresql"
	repoPosts "patreon/internal/app/repository/posts"
	repPostsPsql "patreon/internal/app/repository/posts/postgresql"
	repoPostsData "patreon/internal/app/repository/posts_data"
	repoPostsDataPsql "patreon/internal/app/repository/posts_data/postgresql"
	repUser "patreon/internal/app/repository/user"
	repUserPsql "patreon/internal/app/repository/user/postgresql"
	"patreon/internal/app/sessions"
	"patreon/internal/app/sessions/repository"

	"github.com/sirupsen/logrus"
)

type RepositoryFactory struct {
	expectedConnections app.ExpectedConnections
	logger              *logrus.Logger
	userRepository      repUser.Repository
	creatorRepository   repCreator.Repository
	sessionRepository   sessions.SessionRepository
	awardsRepository    repoAwrds.Repository
	postsRepository     repoPosts.Repository
	likesRepository     repoLikes.Repository
	PostDataRepository  repoPostsData.Repository
	csrfRepository        repCsrf.Repository
	accessRepository      repository_access.Repository
	subscribersRepository repository_subscribers.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections app.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() repUser.Repository {
	if f.userRepository == nil {
		f.userRepository = repUserPsql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}

func (f *RepositoryFactory) GetCreatorRepository() repCreator.Repository {
	if f.creatorRepository == nil {
		f.creatorRepository = repCreatorPsql.NewCreatorRepository(f.expectedConnections.SqlConnection)
	}
	return f.creatorRepository
}

func (f *RepositoryFactory) GetCsrfRepository() repCsrf.Repository {
	if f.csrfRepository == nil {
		f.csrfRepository = repCsrf.NewJwtRepository()
	}
	return f.csrfRepository
}

func (f *RepositoryFactory) GetSessionRepository() sessions.SessionRepository {
	if f.sessionRepository == nil {
		f.sessionRepository = repository.NewRedisRepository(f.expectedConnections.SessionRedisPool, f.logger)
	}
	return f.sessionRepository
}

func (f *RepositoryFactory) GetAwardsRepository() repoAwrds.Repository {
	if f.awardsRepository == nil {
		f.awardsRepository = repAwardsPsql.NewAwardsRepository(f.expectedConnections.SqlConnection)
	}
	return f.awardsRepository
}
func (f *RepositoryFactory) GetAccessRepository() repository_access.Repository {
	if f.accessRepository == nil {
		f.accessRepository = repository_access.NewRedisRepository(f.expectedConnections.SessionRedisPool, f.logger)
	}
	return f.accessRepository
}
func (f *RepositoryFactory) GetSubscribersRepository() repository_subscribers.Repository {
	if f.subscribersRepository == nil {
		f.subscribersRepository = repository_subscribers.NewSubscribersRepository(f.expectedConnections.SqlConnection)
	}
	return f.subscribersRepository
}

func (f *RepositoryFactory) GetPostsRepository() repoPosts.Repository {
	if f.postsRepository == nil {
		f.postsRepository = repPostsPsql.NewPostsRepository(f.expectedConnections.SqlConnection)
	}
	return f.postsRepository
}

func (f *RepositoryFactory) GetLikesRepository() repoLikes.Repository {
	if f.likesRepository == nil {
		f.likesRepository = repLikesPsql.NewLikesRepository(f.expectedConnections.SqlConnection)
	}
	return f.likesRepository
}

func (f *RepositoryFactory) GetPostsDataRepository() repoPostsData.Repository {
	if f.PostDataRepository == nil {
		f.PostDataRepository = repoPostsDataPsql.NewPostsDataRepository(f.expectedConnections.SqlConnection)
	}
	return f.PostDataRepository
}
