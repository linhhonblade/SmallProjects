package usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"social_todo_app_go/common"
)

type TokenParser interface {
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}
type introspectUsecase struct {
	userQueryRepo    UserQueryRepository
	sessionQueryRepo SessionQueryRepository
	tokenParser      TokenParser
}

func NewIntrospectUC(userQueryRepo UserQueryRepository, sessionQueryRepo SessionQueryRepository, tokenParser TokenParser) *introspectUsecase {
	return &introspectUsecase{userQueryRepo: userQueryRepo, sessionQueryRepo: sessionQueryRepo, tokenParser: tokenParser}
}

func (uc *introspectUsecase) IntrospectToken(ctx context.Context, tokenString string) (common.Requester, error) {
	claims, err := uc.tokenParser.ParseToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}
	userId := uuid.MustParse(claims.Subject)
	sessionId := uuid.MustParse(claims.ID)
	_, err = uc.sessionQueryRepo.Find(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	user, err := uc.userQueryRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, errors.New("user is banned")
	}
	return common.NewRequester(
		userId,
		sessionId,
		user.FirstName(),
		user.LastName(),
		user.Role().String(),
		user.Status(),
	), nil
}

func (uc *introspectUsecase) IntrospectCookie(ctx context.Context, tokenString string) (common.Requester, error) {
	claims, err := uc.tokenParser.ParseToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}
	userId := uuid.MustParse(claims.Subject)
	sessionId := uuid.MustParse(claims.ID)
	_, err = uc.sessionQueryRepo.Find(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	user, err := uc.userQueryRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, errors.New("user is banned")
	}
	return common.NewRequester(
		userId,
		sessionId,
		user.FirstName(),
		user.LastName(),
		user.Role().String(),
		user.Status(),
	), nil
}
