package services

import (
	"context"
	"errors"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type TaskService interface {
	CreateTask(ctx context.Context, newTask *generated.Task) (*generated.Task, error)
	GetTasks(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Task, int, error)
	GetTaskByID(ctx context.Context, id int) (*generated.Task, error)
}

type taskService struct {
	client *generated.TaskClient
}

func NewTaskService(client *generated.TaskClient) TaskService {
	return &taskService{client: client}
}

func (s *taskService) CreateTask(ctx context.Context, newTask *generated.Task) (*generated.Task, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	task, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetNillableTaskMilestoneID(newTask.TaskMilestoneID).
		SetProjectID(newTask.ProjectID).
		SetTitle(newTask.Title).
		SetDescription(newTask.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTasks(ctx context.Context, queryParams *render.QueryParams) ([]*generated.Task, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)

	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Task{
		task.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := task.Or(
			task.Title(queryParams.Query),
			task.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var tasks []*generated.Task
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		tasks, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				// TODO: Add others ways to sort
				task.ByCreatedAt(
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

	return tasks, totalFiltered, nil
}

func (s *taskService) GetTaskByID(ctx context.Context, id int) (*generated.Task, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	task, err := s.client.
		Query().
		Where(
			task.AccountIDEQ(int(claims["account_id"].(float64))),
			task.IDEQ(id),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return task, nil
}
