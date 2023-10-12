package services

import (
	"context"
	"errors"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/team"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type TeamService interface {
	CreateTeam(ctx context.Context, newTeam *generated.Team) (*generated.Team, error)
	GetTeams(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Team, int, error)
	GetTeamByID(ctx context.Context, id int) (*generated.Team, error)
	UpdateTeam(ctx context.Context, id int, updatedTeam *generated.Team) (*generated.Team, error)
}

type teamService struct {
	client *generated.TeamClient
}

func NewTeamService(client *generated.TeamClient) TeamService {
	return &teamService{client: client}
}

func (s *teamService) CreateTeam(ctx context.Context, newTeam *generated.Team) (*generated.Team, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	team, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetName(newTeam.Name).
		SetNillableStatus(newTeam.Status).
		SetNillableDescription(newTeam.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *teamService) GetTeams(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Team, int, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Team{
		team.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := team.Or(
			team.NameContains(queryParams.Query),
			team.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var teams []*generated.Team
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		teams, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				team.ByCreatedAt(
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

	return teams, totalFiltered, nil
}

func (s *teamService) GetTeamByID(ctx context.Context, id int) (*generated.Team, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	teams, err := s.client.
		Query().
		Where(
			team.IDEQ(id),
			team.AccountIDEQ(int(claims["account_id"].(float64))),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *teamService) UpdateTeam(ctx context.Context, id int, updatedTeam *generated.Team) (*generated.Team, error) {

	team, err := s.client.UpdateOneID(id).
		SetName(updatedTeam.Name).
		SetNillableStatus(updatedTeam.Status).
		SetNillableDescription(updatedTeam.Description).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return team, nil
}
