package services

import (
	"context"
	"errors"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/projectserviceprovider"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type ProjectServiceProviderService interface {
	GetProjectServiceProviders(ctx context.Context, projectID int, queryParams *render.QueryParams) ([]*generated.ProjectServiceProvider, int, error)
	AddServiceProviderToProject(ctx context.Context, projectID int, serviceProviderID int) (*generated.ProjectServiceProvider, error)
	GetExisting(ctx context.Context, projectID int, serviceProviderID int) (*generated.ProjectServiceProvider, error)
}

type projectServiceProviderService struct {
	client *generated.ProjectServiceProviderClient
}

func NewProjectServiceProviderService(client *generated.ProjectServiceProviderClient) ProjectServiceProviderService {
	return &projectServiceProviderService{client: client}
}

func (ps *projectServiceProviderService) GetProjectServiceProviders(ctx context.Context, projectID int, queryParams *render.QueryParams) ([]*generated.ProjectServiceProvider, int, error) {

	baseFilters := []predicate.ProjectServiceProvider{
		projectserviceprovider.ProjectProjectIDEQ(projectID),
	}

	var wg sync.WaitGroup
	var items []*generated.ProjectServiceProvider
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		items, err1 = ps.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				projectserviceprovider.ByCreatedAt(
					sql.OrderDesc(),
				),
			).
			WithServiceProvider().
			Offset((queryParams.Page - 1) * queryParams.Limit).
			Limit(queryParams.Limit).
			All(ctx)
	}()

	go func() {
		defer wg.Done()
		totalFiltered, err2 = ps.client.
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

	return items, totalFiltered, nil
}

func (ps *projectServiceProviderService) AddServiceProviderToProject(ctx context.Context, projectID int, serviceProviderID int) (*generated.ProjectServiceProvider, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	newProvider, err := ps.client.
		Create().
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetProjectID(projectID).
		SetServiceProviderID(serviceProviderID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return newProvider, nil
}

func (ps *projectServiceProviderService) GetExisting(ctx context.Context, projectID int, serviceProviderID int) (*generated.ProjectServiceProvider, error) {

	provider, err := ps.client.
		Query().
		Where(
			projectserviceprovider.And(
				projectserviceprovider.ProjectServiceProviderID(serviceProviderID),
				projectserviceprovider.ProjectProjectIDEQ(projectID),
			),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return provider, nil
}
