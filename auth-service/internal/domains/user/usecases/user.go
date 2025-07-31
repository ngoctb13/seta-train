package usecases

import (
	"context"
	"regexp"
	"strings"

	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/shared-modules/model"
)

const (
	ManagerRole = "MANAGER"
	MemberRole  = "MEMBER"
)

type User struct {
	userRepo repos.IUserRepo
}

func NewUser(userRepo repos.IUserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) CreateUser(ctx context.Context, user *model.User) error {
	if user == nil {
		return ErrUserCannotBeNil
	}

	if strings.TrimSpace(user.Username) == "" {
		return ErrUsernameRequired
	}
	if len(user.Username) < 3 || len(user.Username) > 50 {
		return ErrUsernameTooShort
	}

	if strings.TrimSpace(user.Email) == "" {
		return ErrEmailRequired
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid
	}

	if strings.TrimSpace(user.PasswordHash) == "" {
		return ErrPasswordRequired
	}

	if user.Role != ManagerRole && user.Role != MemberRole {
		return ErrInvalidRole
	}

	existingUser, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return ErrUsernameAlreadyExists
	}

	existingUser, err = u.userRepo.GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyExists
	}

	return u.userRepo.CreateUser(ctx, user)
}

func (u *User) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *User) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.userRepo.GetUserByEmail(ctx, email)
}

func (u *User) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return u.userRepo.GetUserByUsername(ctx, username)
}

func (u *User) AssignRole(ctx context.Context, userID string, role string) error {
	if userID == "" {
		return ErrUserIDRequired
	}
	if role != ManagerRole && role != MemberRole {
		return ErrInvalidRole
	}
	return u.userRepo.UpdateUserRole(ctx, userID, role)
}
