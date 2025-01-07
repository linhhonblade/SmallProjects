package usecase

import (
	"context"
	"errors"
	"social_todo_app_go/module/user/domain"
	"testing"
)

type mockHasher struct{}

func (mockHasher) RandomStr(length int) (string, error) {
	return "abcd", nil
}

func (mockHasher) HashPassword(salt, password string) (string, error) {
	return "mailovemisa", nil
}

type mockUserRepo struct{}

func (mockUserRepo) Create(ctx context.Context, data *domain.User) error {
	return nil
}

func (mockUserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email == "existed@test.com" {
		return &domain.User{}, nil
	}

	if email == "error@test.com" {
		return nil, errors.New("cannot get record")
	}
	return nil, nil
}

func TestUseCase_Register(t *testing.T) {
	uc := NewUseCase(mockUserRepo{}, mockHasher{})

	type testData struct {
		Input    EmailPasswordRegistrationDTO
		Expected error
	}

	table := []testData{
		{
			Input: EmailPasswordRegistrationDTO{
				FirstName: "Mai",
				LastName:  "Pham",
				Email:     "existed@test.com",
				Password:  "1234",
			},
			Expected: domain.ErrEmailHasExisted},
		{
			Input: EmailPasswordRegistrationDTO{
				FirstName: "Mai",
				LastName:  "Pham",
				Email:     "error@test.com",
				Password:  "1234",
			},
			Expected: errors.New("cannot get record"),
		},
	}

	for i := range table {
		actualErr := uc.Register(context.Background(), table[i].Input)
		if actualErr.Error() != table[i].Expected.Error() {
			t.Errorf("Expected error: %v, got: %v", table[i].Expected, actualErr)
		}
	}
}
