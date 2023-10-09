package services

import (
	"context"
	"errors"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/issue"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type IssueService interface {
	CreateIssue(ctx context.Context, newIssue *generated.Issue) (*generated.Issue, error)
	GetIssues(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Issue, int, error)
	GetIssueByID(ctx context.Context, id int) (*generated.Issue, error)
}

type issueService struct {
	client *generated.IssueClient
}

// TODO: complete the fiields
func NewIssueService(client *generated.IssueClient) IssueService {
	return &issueService{client: client}
}

func (s *issueService) CreateIssue(ctx context.Context, newIssue *generated.Issue) (*generated.Issue, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	issue, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetProjectID(newIssue.ProjectID).
		SetTitle(newIssue.Title).
		SetDescription(newIssue.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return issue, nil
}

func (s *issueService) GetIssues(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Issue, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)

	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Issue{
		issue.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := issue.Or(
			issue.Title(queryParams.Query),
			issue.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var issues []*generated.Issue
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		issues, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				issue.ByCreatedAt(
					sql.OrderDesc(),
				),
			).
			Offset((queryParams.Page - 1) * queryParams.Limit).
			Limit(queryParams.Limit).
			WithUser().
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

	return issues, totalFiltered, nil
}

func (s *issueService) GetIssueByID(ctx context.Context, id int) (*generated.Issue, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	issue, err := s.client.
		Query().
		Where(
			issue.AccountIDEQ(int(claims["account_id"].(float64))),
			issue.IDEQ(id),
		).
		WithUser().
		First(ctx)

	if err != nil {
		return nil, err
	}

	return issue, nil
}
