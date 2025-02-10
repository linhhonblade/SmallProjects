package builder

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"social_todo_app_go/common"
	"social_todo_app_go/module/attachment"
	"social_todo_app_go/module/user/domain"
	"social_todo_app_go/module/user/infras/repository"
	"social_todo_app_go/module/user/usecase"
)

type simpleBuilder struct {
	db *gorm.DB
	tp usecase.TokenProvider
}

func NewSimpleBuilder(db *gorm.DB, tp usecase.TokenProvider) simpleBuilder {
	return simpleBuilder{db: db, tp: tp}
}

func (s simpleBuilder) BuildUserQueryRepo() usecase.UserQueryRepository {
	return repository.NewUserRepo(s.db)
}

func (s simpleBuilder) BuildUserCmdRepo() usecase.UserCommandRepository {
	return repository.NewUserRepo(s.db)
}

func (s simpleBuilder) BuildHasher() usecase.Hasher {
	return &common.Hasher{}
}

func (s simpleBuilder) BuildTokenProvider() usecase.TokenProvider {
	return s.tp
}

func (s simpleBuilder) BuildSessionQueryRepo() usecase.SessionQueryRepository {
	return repository.NewUserSessionPostgresRepo(s.db)
}

func (s simpleBuilder) BuildSessionCmdRepo() usecase.SessionCommandRepository {
	return repository.NewUserSessionPostgresRepo(s.db)
}

func (s simpleBuilder) BuildSessionRepo() usecase.SessionRepository {
	return repository.NewUserSessionPostgresRepo(s.db)
}

func (s simpleBuilder) BuildUserRepo() usecase.UserRepository {
	return repository.NewUserRepo(s.db)
}

func (s simpleBuilder) BuildAttachmentRepo() usecase.AttachmentRepository {
	return attachment.NewRepo(s.db)
}

type complexBuilder struct {
	simpleBuilder
}

type userCacheRepo struct {
	realRepo usecase.UserQueryRepository
	cache    map[string]*domain.User
}

func NewComplexBuilder(simpleBuilder simpleBuilder) complexBuilder {
	return complexBuilder{
		simpleBuilder: simpleBuilder,
	}
}

func (c userCacheRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if user, ok := c.cache[email]; ok {
		return user, nil
	}
	user, err := c.realRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	c.cache[email] = user
	return user, nil
}

func (c userCacheRepo) FindByID(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	if user, ok := c.cache[userId.String()]; ok {
		return user, nil
	}
	user, err := c.realRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	c.cache[userId.String()] = user
	return user, nil
}

func (cb complexBuilder) BuildUserQueryRepo() usecase.UserQueryRepository {
	return userCacheRepo{realRepo: cb.simpleBuilder.BuildUserQueryRepo(), cache: make(map[string]*domain.User)}
}
