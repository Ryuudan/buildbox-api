package services

import (
	"context"
	"errors"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/milestone"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type MilestoneService interface {
	CreateMilestone(ctx context.Context, newMilestone *generated.Milestone) (*generated.Milestone, error)
	GetMilestones(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Milestone, int, error)
	GetMilestoneByID(ctx context.Context, id int) (*generated.Milestone, error)
}

type milestoneService struct {
	client *generated.MilestoneClient
}

func NewMilestoneService(client *generated.MilestoneClient) MilestoneService {
	return &milestoneService{client: client}
}

func (s *milestoneService) CreateMilestone(ctx context.Context, newMilestone *generated.Milestone) (*generated.Milestone, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	milestone, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetProjectID(newMilestone.ProjectID).
		SetEndDate(newMilestone.EndDate).
		SetTitle(newMilestone.Title).
		SetDescription(newMilestone.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return milestone, nil
}

func (s *milestoneService) GetMilestoneByID(ctx context.Context, id int) (*generated.Milestone, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	milestone, err := s.client.Query().
		Where(
			milestone.AccountIDEQ(int(claims["account_id"].(float64))),
			milestone.IDEQ(id),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return milestone, nil
}

func (s *milestoneService) GetMilestones(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Milestone, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Milestone{
		milestone.AccountIDEQ(int(claims["account_id"].(float64))),
		milestone.DeletedEQ(false),
	}

	// baseOrders := getProjectOrders(filters.Order)

	if queryParams.Query != "" {
		searchFilter := milestone.Or(
			milestone.TitleContains(queryParams.Query),
			milestone.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var milestones []*generated.Milestone
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		milestones, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				milestone.ByCreatedAt(
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

	return milestones, totalFiltered, nil
}
