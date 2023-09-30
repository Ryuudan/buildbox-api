package services

import (
	"context"
	"errors"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
)

type UserService interface {
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

func (s *userService) RegisterUser(ctx context.Context, newUser *generated.User) (*generated.User, error) {

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

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*generated.User, error) {
	user, err := s.client.Query().Where(
		user.EmailEQ(email),
	).WithAccount().First(ctx)

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
