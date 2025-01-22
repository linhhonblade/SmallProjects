package usecase

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/user/domain"
	"time"
)

type loginEmailPasswordUC struct {
	userRepo      UserQueryRepository
	tokenProvider TokenProvider
	sessionRepo   SessionCommandRepository
	hasher        Hasher
}

func NewLoginEmailPasswordUC(userRepo UserQueryRepository, tokenProvider TokenProvider, sessionRepo SessionCommandRepository, hasher Hasher) *loginEmailPasswordUC {
	return &loginEmailPasswordUC{userRepo: userRepo, tokenProvider: tokenProvider, sessionRepo: sessionRepo, hasher: hasher}
}

func (uc *loginEmailPasswordUC) LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error) {
	// 1. Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	// 2. Hash and compare password
	if ok := uc.hasher.CompareHashPassword(user.Password(), user.Salt(), dto.Password); !ok {
		return nil, domain.InvalidEmailPassword
	}

	userId := user.Id()
	sessionId := common.GenUUID()

	// 3. Generate JWT
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, err
	}

	// 4. Insert Session into DB
	refreshToken, _ := uc.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpiredInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenRefreshInSeconds()))
	session := domain.NewUserSession(sessionId, userId, refreshToken, tokenExpAt, refreshExpAt)
	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	// 5.Return token response dto

	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.TokenExpiredInSeconds(),
		RefreshToken:      refreshToken,
		RefreshTokenExpIn: uc.tokenProvider.TokenRefreshInSeconds()}, nil
}
