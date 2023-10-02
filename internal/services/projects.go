package services

import (
	"context"
	"errors"
	"log"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// ProjectService defines methods for project management.
type ProjectService interface {
	CreateProject(ctx context.Context, newProject *generated.Project) (*generated.Project, error)
	GetProjects(ctx context.Context, queryParams *render.QueryParams, filters models.Filters) ([]*generated.Project, int, error)
	UpdateProjectByID(ctx context.Context, id int, newPayload *generated.Project) (*generated.Project, error)
	GetProjectByID(ctx context.Context, id int) (*generated.Project, error)
	GetProjectByUUID(ctx context.Context, uuid uuid.UUID) (*generated.Project, error)
	DeleteManyProjectsByID(ctx context.Context, ids []int) ([]*generated.Project, error)
	DeleteProjectByID(ctx context.Context, id int) (*generated.Project, error)
}

// EntProjectService is an implementation of ProjectService that uses the Ent client.
type projectService struct {
	client *generated.ProjectClient
}

// NewEntProjectService creates a new instance of EntProjectService.
func NewProjectService(client *generated.ProjectClient) ProjectService {
	return &projectService{client: client}
}

// CreateProject creates a new project.
func (s *projectService) CreateProject(ctx context.Context, newProject *generated.Project) (*generated.Project, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	project, err := s.client.Create().
		SetAccountID(int(claims["account_id"].(float64))).
		SetCreatedBy(int(claims["user_id"].(float64))).
		SetName(newProject.Name).
		SetNillableClientID(newProject.ClientID).
		SetNillableManagerID(newProject.ManagerID).
		SetNillableStatus(newProject.Status).
		SetNillableStartDate(newProject.StartDate).
		SetNillableEndDate(newProject.EndDate).
		SetNillableBudget(newProject.Budget).
		SetNillableLocation(newProject.Location).
		SetNillableDescription(newProject.Description).
		SetNillableNotes(newProject.Notes).
		SetDeleted(newProject.Deleted).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return project, nil
}

// func GetSortFields(fields []models.SortFields, target interface) {

// }

func (s *projectService) GetProjects(ctx context.Context, queryParams *render.QueryParams, filters models.Filters) ([]*generated.Project, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	// Base Filters
	baseFilters := []predicate.Project{
		project.AccountIDEQ(int(claims["account_id"].(float64))),
		project.DeletedEQ(false),
	}

	utils.ConsoleLog(filters)
	if queryParams.Query != "" {
		orCondition := project.Or(
			project.NameContains(queryParams.Query),
			project.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, orCondition)
	}

	var wg sync.WaitGroup
	var projects []*generated.Project
	var totalProjects int
	var err1, err2 error

	wg.Add(2)
	// Get the projects
	go func() {

		baseOrders := []project.OrderOption{}

		if len(filters.Order) == 0 {
			baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderDesc()))
		}

		for _, sortCriteria := range filters.Order {
			switch sortCriteria.Field {
			case "name":
				if sortCriteria.Direction == "asc" {
					baseOrders = append(baseOrders, project.ByName(sql.OrderAsc()))
				} else if sortCriteria.Direction == "desc" {
					baseOrders = append(baseOrders, project.ByName(sql.OrderDesc()))
				}
			case "budget":
				if sortCriteria.Direction == "asc" {
					baseOrders = append(baseOrders, project.ByBudget(sql.OrderAsc()))
				} else if sortCriteria.Direction == "desc" {
					baseOrders = append(baseOrders, project.ByBudget(sql.OrderDesc()))
				}
			case "created_at":
				if sortCriteria.Direction == "asc" {
					baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderAsc()))
				} else if sortCriteria.Direction == "desc" {
					baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderDesc()))
				}
			default:
				baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderDesc()))
				log.Println("Unknown Field is passed as order ignore it")
				// Handle unknown fields or provide a default behavior
				// For example, you can log an error or ignore unknown fields
			}

		}

		defer wg.Done()
		projects, err1 = s.client.Query().
			Where(baseFilters...).
			Offset((queryParams.Page - 1) * queryParams.Limit).
			Limit(queryParams.Limit).
			Order(
				baseOrders...,
			).
			All(ctx)
	}()

	// Get the total projects with the current filters
	go func() {
		defer wg.Done()
		totalProjects, err2 = s.client.Query().Where(baseFilters...).Aggregate(generated.Count()).Int(ctx)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, 0, err1
	}

	if err2 != nil {
		return nil, 0, err2
	}

	return projects, totalProjects, nil
}

// UpdateProject updates an existing project.
func (s *projectService) UpdateProjectByID(ctx context.Context, id int, newPayload *generated.Project) (*generated.Project, error) {
	project, err := s.client.UpdateOneID(id).
		SetName(newPayload.Name).
		SetStartDate(*newPayload.StartDate).
		SetStatus(*newPayload.Status).
		SetEndDate(*newPayload.EndDate).
		SetBudget(*newPayload.Budget).
		SetLocation(*newPayload.Location).
		SetDescription(*newPayload.Description).
		SetNotes(*newPayload.Notes).
		SetManagerID(*newPayload.ManagerID).
		SetDeleted(newPayload.Deleted).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProjectById sets the 'deleted' field to true for a project.
func (s *projectService) DeleteProjectByID(ctx context.Context, id int) (*generated.Project, error) {
	deleted := true

	project, err := s.client.UpdateOneID(id).
		SetNillableDeleted(&deleted).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteManyProjects sets the 'deleted' field to true for multiple projects by their IDs.
func (s *projectService) DeleteManyProjectsByID(ctx context.Context, ids []int) ([]*generated.Project, error) {
	deleted := true

	// Create a slice to store the updated projects.
	var deletedProjects []*generated.Project

	// Iterate through the project IDs and update each project.
	for _, id := range ids {
		project, err := s.client.UpdateOneID(id).
			SetNillableDeleted(&deleted).
			Save(ctx)

		if err != nil {
			return nil, err
		}

		deletedProjects = append(deletedProjects, project)
	}

	return deletedProjects, nil
}

// GetProjectByID retrieves a project by its ID.
func (s *projectService) GetProjectByID(ctx context.Context, id int) (*generated.Project, error) {
	project, err := s.client.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("project not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}
	return project, nil
}

// GetProjectByUUID retrieves a project by its UUID.
func (s *projectService) GetProjectByUUID(ctx context.Context, uuid uuid.UUID) (*generated.Project, error) {
	project, err := s.client.Query().
		Where(project.UUIDEQ(uuid)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// project.ByBudget(
// 	sql.OrderDesc(),
// ),
// project.ByName(
// 	sql.OrderDesc(),
// ),
// project.ByStatus(
// 	sql.OrderDesc(),
// ),
// project.ByLocation(
// 	sql.OrderDesc(),
// ),
// project.ByBudget(
// 	sql.OrderDesc(),
// ),
// project.ByDescription(
// 	sql.OrderDesc(),
// ),
// project.ByNotes(
// 	sql.OrderDesc(),
// ),
// project.ByStartDate(
// 	sql.OrderDesc(),
// ),
// project.ByEndDate(
// 	sql.OrderDesc(),
// ),
// project.ByDeleted(
// 	sql.OrderDesc(),
// ),
