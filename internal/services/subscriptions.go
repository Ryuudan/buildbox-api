package services

import (
	"context"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, newSub *generated.Subscription) (*generated.Subscription, error)
	GetActiveSubscriptionByAccountID(ctx context.Context, accountID int) (*generated.Subscription, error)
	UpdateSubscriptionStatusByID(ctx context.Context, id int, status subscription.Status) (*generated.Subscription, error)
}

type subscriptionService struct {
	client *generated.SubscriptionClient
}

func NewSubscriptionService(client *generated.SubscriptionClient) SubscriptionService {
	return &subscriptionService{client: client}
}

func (s *subscriptionService) GetActiveSubscriptionByAccountID(ctx context.Context, accountID int) (*generated.Subscription, error) {
	subscription, err := s.client.Query().Where(
		subscription.AccountIDEQ(accountID),
		subscription.StatusEQ("active"),
	).WithPlan().First(ctx)

	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (s *subscriptionService) UpdateSubscriptionStatusByID(ctx context.Context, id int, status subscription.Status) (*generated.Subscription, error) {
	updated, err := s.client.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *subscriptionService) CreateSubscription(ctx context.Context, newSub *generated.Subscription) (*generated.Subscription, error) {
	subscription, err := s.client.Create().
		SetAccountID(newSub.AccountID).
		SetPlanID(newSub.PlanID).
		SetStartDate(newSub.StartDate).
		SetEndDate(newSub.EndDate).
		SetDiscount(newSub.Discount).
		SetStatus(newSub.Status).
		SetBillingCycle(newSub.BillingCycle).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}
