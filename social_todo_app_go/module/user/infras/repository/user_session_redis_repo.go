package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"social_todo_app_go/module/user/domain"
	"time"
)

type sessionRedisRepo struct {
	client *redis.Client
}

func NewSessionRedisRepo(client *redis.Client) *sessionRedisRepo {
	return &sessionRedisRepo{
		client: client,
	}
}

func (r *sessionRedisRepo) Create(ctx context.Context, data *domain.UserSession) error {
	dto := userSessionDto{
		Id:           data.Id(),
		UserId:       data.UserId(),
		RefreshToken: data.RefreshToken(),
		AccessExpAt:  data.AccessExpAt(),
		RefreshExpAt: data.RefreshExpAt(),
	}
	key := "session:" + dto.Id.String()
	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":             dto.Id.String(),
		"user_id":        dto.UserId.String(),
		"refresh_token":  dto.RefreshToken,
		"access_exp_at":  dto.AccessExpAt.Format("2006-01-02 15:04:05"),
		"refresh_exp_at": dto.RefreshExpAt.Format("2006-01-02 15:04:05"),
	}).Err()
	if err != nil {
		return err
	}
	// Set expiration for the session
	// TODO: Should be slightly longer than access token expiration
	r.client.Expire(context.Background(), key, 15*time.Minute)
	// Lưu refresh token vào Redis dùng SET (key riêng)
	// TODO: Should be slightly longer than refresh token expiration
	err = r.client.Set(context.Background(), "refresh:"+dto.RefreshToken, dto.Id.String(), 7*24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *sessionRedisRepo) Delete(ctx context.Context, sessionId uuid.UUID) error {
	if err := r.client.Del(ctx, "session:"+sessionId.String()).Err(); err != nil {
		return err
	}
	return nil
}

func (r *sessionRedisRepo) Find(ctx context.Context, sessionId uuid.UUID) (*domain.UserSession, error) {
	key := "session:" + sessionId.String()
	data, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil // Not found
	}

	id := uuid.MustParse(data["id"])
	userId := uuid.MustParse(data["user_id"])
	accessExpAt, _ := time.Parse("2006-01-02 15:04:05", data["access_exp_at"])
	refreshExpAt, _ := time.Parse("2006-01-02 15:04:05", data["refresh_exp_at"])

	return domain.NewUserSession(
		id,
		userId,
		data["refresh_token"],
		accessExpAt,
		refreshExpAt,
	), nil
}

func (r *sessionRedisRepo) FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.UserSession, error) {
	key := "refresh:" + refreshToken
	sessionId, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	session_key := "session:" + sessionId
	data, err := r.client.HGetAll(ctx, session_key).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil // Not found
	}

	id := uuid.MustParse(data["id"])
	userId := uuid.MustParse(data["user_id"])
	accessExpAt, _ := time.Parse("2006-01-02 15:04:05", data["access_exp_at"])
	refreshExpAt, _ := time.Parse("2006-01-02 15:04:05", data["refresh_exp_at"])

	return domain.NewUserSession(
		id,
		userId,
		data["refresh_token"],
		accessExpAt,
		refreshExpAt,
	), nil
}
