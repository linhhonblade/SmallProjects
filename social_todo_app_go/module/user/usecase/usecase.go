package usecase

import (
	"context"
	"social_todo_app_go/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error
	LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashPassword, salt, password string) bool
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpiredInSeconds() int
	TokenRefreshInSeconds() int
}

type useCase struct {
	*registerUC
	*loginEmailPasswordUC
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCmdRepo() UserCommandRepository
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCmdRepo() SessionCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
}

func NewUCWithBuilder(b Builder) UseCase {
	return &useCase{
		registerUC:           NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildHasher()),
		loginEmailPasswordUC: NewLoginEmailPasswordUC(b.BuildUserQueryRepo(), b.BuildTokenProvider(), b.BuildSessionCmdRepo(), b.BuildHasher()),
	}
}

func NewUseCase(repo UserRepository, hasher Hasher, tokenProvider TokenProvider, sessionRepo SessionRepository) UseCase {
	return &useCase{
		registerUC:           NewRegisterUC(repo, repo, hasher),
		loginEmailPasswordUC: NewLoginEmailPasswordUC(repo, tokenProvider, sessionRepo, hasher)}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}

type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, data *domain.User) error
}

type SessionRepository interface {
	SessionCommandRepository
}

type SessionCommandRepository interface {
	Create(ctx context.Context, session *domain.UserSession) error
}

type SessionQueryRepository interface{}
