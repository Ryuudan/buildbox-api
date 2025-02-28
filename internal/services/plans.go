package services

import (
	"context"
	"strings"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/plan"
	"github.com/Pyakz/buildbox-api/internal/models"
)

type PlanService interface {
	CreatePlan(ctx context.Context, plan *generated.Plan) (*generated.Plan, error)
	GetPlans(ctx context.Context) ([]*generated.Plan, error)
	GetPlanByName(ctx context.Context, name string) (*generated.Plan, error)
	GetPlanByID(ctx context.Context, id int) (*generated.Plan, error)
	UpdatePlan(ctx context.Context, id int, updated models.UpdatePlanRequest) (*generated.Plan, error)
}

type planService struct {
	client *generated.PlanClient
}

func NewPlanService(client *generated.PlanClient) PlanService {
	return &planService{client: client}
}

func (p *planService) CreatePlan(ctx context.Context, plan *generated.Plan) (*generated.Plan, error) {

	plan, err := p.client.
		Create().
		SetName(strings.ToLower(plan.Name)).
		SetPrice(*plan.Price).
		SetDescription(plan.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p *planService) GetPlans(ctx context.Context) ([]*generated.Plan, error) {
	plans, err := p.client.Query().Select().All(ctx)

	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (p *planService) GetPlanByName(ctx context.Context, name string) (*generated.Plan, error) {

	plan, err := p.client.Query().Where(
		plan.NameEQ(name),
	).First(ctx)

	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p *planService) GetPlanByID(ctx context.Context, id int) (*generated.Plan, error) {
	plan, err := p.client.Query().Where(
		plan.IDEQ(id),
	).First(ctx)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p *planService) UpdatePlan(ctx context.Context, id int, updated models.UpdatePlanRequest) (*generated.Plan, error) {

	plan, err := p.client.UpdateOneID(id).
		SetNillableName(updated.Name).
		SetNillableDescription(updated.Description).
		SetNillablePrice(updated.Price).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return plan, nil
}
