package services

import (
	"context"
	"errors"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, newSub *generated.Subscription) (*generated.Subscription, error)
	GetSubscriptionsByAccountID(ctx context.Context, accountID int) ([]*generated.Subscription, error)
	GetActiveSubscriptionsByAccountID(ctx context.Context, accountID int) ([]*generated.Subscription, error)
	UpdateSubscriptionStatusByID(ctx context.Context, id int, status subscription.Status) (*generated.Subscription, error)
}

type subscriptionService struct {
	client *generated.SubscriptionClient
}

func NewSubscriptionService(client *generated.SubscriptionClient) SubscriptionService {
	return &subscriptionService{client: client}
}

func (s *subscriptionService) GetSubscriptionsByAccountID(ctx context.Context, accountID int) ([]*generated.Subscription, error) {

	subscriptions, err := s.client.Query().Where(
		subscription.AccountIDEQ(accountID),
	).All(ctx)

	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("subscription not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}

	return subscriptions, nil
}

func (s *subscriptionService) GetActiveSubscriptionsByAccountID(ctx context.Context, accountID int) ([]*generated.Subscription, error) {

	subscriptions, err := s.client.Query().Where(
		subscription.AccountIDEQ(accountID),
		subscription.StatusEQ("active"),
	).All(ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("subscription not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}

	return subscriptions, nil
}

func (s *subscriptionService) UpdateSubscriptionStatusByID(ctx context.Context, id int, status subscription.Status) (*generated.Subscription, error) {
	updated, err := s.client.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("subscription not found")
		}
		return nil, errors.New("something went wrong, please try again later")
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
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}
