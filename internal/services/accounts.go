package services

import (
	"context"
	"errors"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/account"

	_ "github.com/lib/pq"
)

type AccountService interface {
	CreateAccount(ctx context.Context, newAccount *generated.Account) (*generated.Account, error)
	GetAccounts(ctx context.Context) ([]*generated.Account, error)
	GetAccountByID(ctx context.Context, id int) (*generated.Account, error)
}

type accountService struct {
	client *generated.AccountClient
}

func NewAccountService(client *generated.AccountClient) AccountService {
	return &accountService{client: client}
}

func (s *accountService) CreateAccount(ctx context.Context, newAccount *generated.Account) (*generated.Account, error) {

	account, err := s.client.Create().
		SetName(newAccount.Name).
		SetEmail(newAccount.Email).
		SetPhoneNumber(newAccount.PhoneNumber).
		Save(ctx)
	if err != nil {

		if generated.IsConstraintError(err) {
			return nil, errors.New("email already exists")
		}

		return nil, errors.New("something went wrong, please try again later")
	}

	return account, nil
}

func (s *accountService) GetAccounts(ctx context.Context) ([]*generated.Account, error) {
	accounts, err := s.client.Query().Select(
		// return all fields
		account.FieldID,
		account.FieldName,
	).All(ctx)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *accountService) GetAccountByID(ctx context.Context, id int) (*generated.Account, error) {
	account, err := s.client.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("account not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}
	return account, nil
}
