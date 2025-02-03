package usecase

import (
	"context"
	"errors"
	"social_todo_app_go/common"
	"social_todo_app_go/module/user/domain"
	"time"
)

type refreshTokenUC struct {
	userRepo      UserRepository
	sessionRepo   SessionRepository
	tokenProvider TokenProvider
	hasher        Hasher
}

func NewRefreshTokenUC(userRepo UserRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) *refreshTokenUC {
	return &refreshTokenUC{
		userRepo:      userRepo,
		sessionRepo:   sessionRepo,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}

func (r *refreshTokenUC) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error) {
	//1. Find session by refresh token
	session, err := r.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	//Check if refresh token is expired
	if session.RefreshExpAt().Before(time.Now().UTC()) {
		return nil, errors.New("refresh token is expired")
	}

	//2. Find user by session user id
	user, err := r.userRepo.FindByID(ctx, session.UserId())
	if err != nil {
		return nil, err
	}
	if user.Status() == "banned" {
		return nil, errors.New("user is banned")
	}
	//3. Generate new jwt access token
	userId := user.Id()
	sessionId := common.GenUUID()
	accessToken, err := r.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, err
	}

	//4. Update session with new access token
	newRefreshToken, _ := r.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(r.tokenProvider.TokenExpiredInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(r.tokenProvider.TokenRefreshInSeconds()))
	newSession := domain.NewUserSession(sessionId, userId, newRefreshToken, tokenExpAt, refreshExpAt)
	if err := r.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	//5. Delete old sessions
	go func() {
		_ = r.sessionRepo.Delete(ctx, session.Id())
	}()

	//6. Return token response dto
	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  r.tokenProvider.TokenExpiredInSeconds(),
		RefreshToken:      newRefreshToken,
		RefreshTokenExpIn: r.tokenProvider.TokenRefreshInSeconds(),
	}, nil
}
