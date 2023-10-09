package services

import (
	"context"
	"errors"
	"sync"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/utils/render"

	_ "github.com/lib/pq"
)

type AccountService interface {
	CreateAccount(ctx context.Context, newAccount *generated.Account) (*generated.Account, error)
	GetAccounts(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Account, int, error)
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

// TODO: Implement pagination
func (s *accountService) GetAccounts(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Account, int, error) {

	var wg sync.WaitGroup
	var accounts []*generated.Account
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		accounts, err1 = s.client.Query().
			Offset((queryParams.Page - 1) * queryParams.Limit).
			Limit(queryParams.Limit).
			All(ctx)
	}()

	go func() {
		defer wg.Done()
		totalFiltered, err2 = s.client.
			Query().
			Aggregate(generated.Count()).Int(ctx)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, 0, err1
	}

	if err2 != nil {
		return nil, 0, err2
	}

	return accounts, totalFiltered, nil
}

func (s *accountService) GetAccountByID(ctx context.Context, id int) (*generated.Account, error) {
	account, err := s.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
