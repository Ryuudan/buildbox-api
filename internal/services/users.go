package services

import (
	"context"
	"errors"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, newUser *generated.User) (*generated.User, error)
	RegisterUser(ctx context.Context, newUser *generated.User) (*generated.User, error)
	GetUserByEmail(ctx context.Context, email string) (*generated.User, error)
	GetUserById(ctx context.Context, id int) (*generated.User, error)
	GetUsersByAccountID(ctx context.Context, accountID int) ([]*generated.User, error)
}

type userService struct {
	client *generated.UserClient
}

func NewUserService(client *generated.UserClient) UserService {
	return &userService{client: client}
}

// This is when after an account is created, we need to create a user for that account.
// notice that this doesn't have context claims to determine with account the user belongs to.
func (s *userService) RegisterUser(ctx context.Context, newUser *generated.User) (*generated.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to generate hashed password")
	}

	newUser.Password = string(password)

	user, err := s.client.Create().
		SetAccountID(newUser.AccountID).
		SetEmail(newUser.Email).
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetPassword(newUser.Password).
		SetNillableMiddleName(newUser.MiddleName).
		SetNillableBirthday(newUser.Birthday).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateProject creates a new project.
func (s *userService) CreateUser(ctx context.Context, newUser *generated.User) (*generated.User, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to generate hashed password")
	}

	newUser.Password = string(password)

	user, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetEmail(newUser.Email).
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetPassword(newUser.Password).
		SetNillableMiddleName(newUser.MiddleName).
		SetNillableBirthday(newUser.Birthday).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*generated.User, error) {
	user, err := s.client.Query().Where(
		user.EmailEQ(email),
	).First(ctx)

	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}

	return user, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (*generated.User, error) {
	user, err := s.client.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}

	return user, nil
}

func (s *userService) GetUsersByAccountID(ctx context.Context, accountID int) ([]*generated.User, error) {
	users, err := s.client.Query().Where(
		user.AccountIDEQ(accountID),
	).All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
