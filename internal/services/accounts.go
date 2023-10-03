package services

import (
	"context"
	"errors"
	"sync"

	"github.com/Pyakz/buildbox-api/ent/generated"

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

		return nil, errors.New(err.Error())
	}

	return account, nil
}

func (s *accountService) GetAccounts(ctx context.Context) ([]*generated.Account, error) {

	var wg sync.WaitGroup
	var accounts []*generated.Account
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		accounts, err1 = s.client.Query().All(ctx)
	}()

	go func() {
		defer wg.Done()
		totalFiltered, err2 = s.client.
			Query().
			Aggregate(generated.Count()).Int(ctx)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	println("Total Filtered: ", totalFiltered)
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
