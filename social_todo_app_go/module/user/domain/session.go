package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserSession struct {
	id           uuid.UUID
	userId       uuid.UUID
	refreshToken string
	accessExpAt  time.Time
	refreshExpAt time.Time
}

func NewUserSession(id uuid.UUID, userId uuid.UUID, refreshToken string, accessExpAt time.Time, refreshExpAt time.Time) *UserSession {
	return &UserSession{id: id, userId: userId, refreshToken: refreshToken, accessExpAt: accessExpAt, refreshExpAt: refreshExpAt}
}

func (s UserSession) Id() uuid.UUID {
	return s.id
}

func (s UserSession) UserId() uuid.UUID {
	return s.userId
}

func (s UserSession) RefreshToken() string {
	return s.refreshToken
}

func (s UserSession) AccessExpAt() time.Time {
	return s.accessExpAt
}

func (s UserSession) RefreshExpAt() time.Time {
	return s.refreshExpAt
}
