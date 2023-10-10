package services

import (
	"context"
	"errors"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type ServiceProviderService interface {
	CreateServiceProvider(ctx context.Context, newServiceProvider *generated.ServiceProvider) (*generated.ServiceProvider, error)
	GetServiceProviderByEmail(ctx context.Context, email string) (*generated.ServiceProvider, error)
	GetServiceProviderByPhone(ctx context.Context, phone string) (*generated.ServiceProvider, error)
	GetServiceProviderByID(ctx context.Context, id int) (*generated.ServiceProvider, error)
	GetServiceProviders(ctx context.Context, queryParams *render.QueryParams) ([]*generated.ServiceProvider, int, error)
}

type serviceProviderService struct {
	client *generated.ServiceProviderClient
}

func NewServiceProviderService(client *generated.ServiceProviderClient) ServiceProviderService {
	return &serviceProviderService{client: client}
}

func (s *serviceProviderService) CreateServiceProvider(ctx context.Context, newServiceProvider *generated.ServiceProvider) (*generated.ServiceProvider, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	service_provider, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetEmail(newServiceProvider.Email).
		SetName(newServiceProvider.Name).
		SetPhoneNumber(newServiceProvider.PhoneNumber).
		SetNillableStatus(newServiceProvider.Status).
		SetNillableDescription(newServiceProvider.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return service_provider, nil
}

func (s *serviceProviderService) GetServiceProviderByEmail(ctx context.Context, email string) (*generated.ServiceProvider, error) {

	service_provider, err := s.client.Query().Where(
		serviceprovider.EmailEQ(email),
	).WithAccount().First(ctx)

	if err != nil {
		return nil, err
	}

	return service_provider, nil
}

func (s *serviceProviderService) GetServiceProviderByPhone(ctx context.Context, phone string) (*generated.ServiceProvider, error) {

	service_provider, err := s.client.Query().Where(
		serviceprovider.PhoneNumber(phone),
	).WithAccount().First(ctx)

	if err != nil {
		return nil, err
	}

	return service_provider, nil
}

func (s *serviceProviderService) GetServiceProviderByID(ctx context.Context, id int) (*generated.ServiceProvider, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	service_provider, err := s.client.
		Query().
		Where(
			serviceprovider.IDEQ(id),
			serviceprovider.AccountIDEQ(int(claims["account_id"].(float64))),
		).
		WithAccount().
		First(ctx)

	if err != nil {
		return nil, err
	}

	return service_provider, nil
}

func (s *serviceProviderService) GetServiceProviders(ctx context.Context, queryParams *render.QueryParams) ([]*generated.ServiceProvider, int, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.ServiceProvider{
		serviceprovider.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := serviceprovider.Or(
			serviceprovider.NameContains(queryParams.Query),
			serviceprovider.EmailContains(queryParams.Query),
			serviceprovider.PhoneNumber(queryParams.Query),
			serviceprovider.Description(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var service_providers []*generated.ServiceProvider
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		service_providers, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				serviceprovider.ByCreatedAt(
					sql.OrderDesc(),
				),
			).
			Offset((queryParams.Page - 1) * queryParams.Limit).
			Limit(queryParams.Limit).
			All(ctx)
	}()

	go func() {
		defer wg.Done()
		totalFiltered, err2 = s.client.
			Query().
			Where(baseFilters...).
			Aggregate(generated.Count()).Int(ctx)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, 0, err1
	}

	if err2 != nil {
		return nil, 0, err2
	}

	return service_providers, totalFiltered, nil

}
