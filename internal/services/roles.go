package services

import (
	"context"
	"errors"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/role"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type RolesService interface {
	CreateRole(ctx context.Context, newRole *generated.Role) (*generated.Role, error)
	GetRoles(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Role, int, error)
	GetRoleByName(ctx context.Context, name string) (*generated.Role, error)
	GetRoleByID(ctx context.Context, id int) (*generated.Role, error)
	UpdateRole(ctx context.Context, id int, updatedRole models.UpdateRolePayload) (*generated.Role, error)
}

type rolesService struct {
	client *generated.RoleClient
}

func NewRolesService(client *generated.RoleClient) RolesService {
	return &rolesService{client: client}
}

func (s *rolesService) CreateRole(ctx context.Context, newRole *generated.Role) (*generated.Role, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	role, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetName(newRole.Name).
		SetDescription(newRole.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *rolesService) GetRoleByID(ctx context.Context, id int) (*generated.Role, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	role, err := s.client.
		Query().
		Where(
			role.AccountIDEQ(int(claims["account_id"].(float64))),
			role.IDEQ(id),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *rolesService) GetRoles(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Role, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)

	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Role{
		role.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := role.Or(
			role.NameContains(queryParams.Query),
			role.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var roles []*generated.Role
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		roles, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				role.ByCreatedAt(
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

	return roles, totalFiltered, nil
}

func (s *rolesService) GetRoleByName(ctx context.Context, name string) (*generated.Role, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	role, err := s.client.
		Query().
		Where(
			role.AccountIDEQ(int(claims["account_id"].(float64))),
			role.NameEQ(name),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *rolesService) UpdateRole(ctx context.Context, id int, updatedRole models.UpdateRolePayload) (*generated.Role, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	role, err := s.client.
		UpdateOneID(id).
		SetNillableName(updatedRole.Name).
		SetNillableDescription(updatedRole.Description).
		SetUpdatedAt(time.Now()).
		Where(
			role.AccountIDEQ(int(claims["account_id"].(float64))),
		).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return role, nil
}
