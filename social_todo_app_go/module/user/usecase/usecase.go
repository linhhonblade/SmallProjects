package usecase

import (
	"context"
	"github.com/google/uuid"
	"social_todo_app_go/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error
	LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error)
	ChangeAvt(ctx context.Context, dto SetSingleImageDTO) error
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

type ChangeAvatar interface {
	ChangeAvt(ctx context.Context, dto SetSingleImageDTO) error
}

type useCase struct {
	*registerUC
	*loginEmailPasswordUC
	*refreshTokenUC
	*changeAvtUC
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCmdRepo() UserCommandRepository
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCmdRepo() SessionCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionRepo() SessionRepository
	BuildUserRepo() UserRepository
	BuildAttachmentRepo() AttachmentRepository
}

func NewUCWithBuilder(b Builder) UseCase {
	return &useCase{
		registerUC:           NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildHasher()),
		loginEmailPasswordUC: NewLoginEmailPasswordUC(b.BuildUserQueryRepo(), b.BuildTokenProvider(), b.BuildSessionCmdRepo(), b.BuildHasher()),
		refreshTokenUC:       NewRefreshTokenUC(b.BuildUserRepo(), b.BuildSessionRepo(), b.BuildTokenProvider(), b.BuildHasher()),
		changeAvtUC:          NewChangeAvtUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildAttachmentRepo()),
	}
}

func NewUseCase(repo UserRepository, hasher Hasher, tokenProvider TokenProvider, sessionRepo SessionRepository) UseCase {
	return &useCase{
		registerUC:           NewRegisterUC(repo, repo, hasher),
		loginEmailPasswordUC: NewLoginEmailPasswordUC(repo, tokenProvider, sessionRepo, hasher),
		refreshTokenUC:       NewRefreshTokenUC(repo, sessionRepo, tokenProvider, hasher)}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}

type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, data *domain.User) error
	Update(ctx context.Context, cond map[string]interface{}, data *domain.UserUpdate) error
}

type SessionRepository interface {
	SessionCommandRepository
	SessionQueryRepository
}

type SessionCommandRepository interface {
	Create(ctx context.Context, session *domain.UserSession) error
	Delete(ctx context.Context, sessionId uuid.UUID) error
}

type SessionQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*domain.UserSession, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.UserSession, error)
}
