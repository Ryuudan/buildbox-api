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

// GetProjects retrieves a list of projects based on query parameters and filters,
// and returns the projects along with the total count of matching projects, in the current accountID
// It applies user-specific filters based on claims from the context,
// supports searching by name or description, sorting, and pagination.
//
// Parameters:
//   - ctx: The request context.
//   - queryParams: Query parameters specifying page, limit, query, and order.
//   - filters: Additional filters to apply.
//
// Returns:
//   - []*generated.Project: List of projects matching the criteria.
//   - int: Total count of matching projects.
//   - error: An error if any occurred during the retrieval.
func (s *projectService) GetProjects(ctx context.Context, queryParams *render.QueryParams, filters models.Filters) ([]*generated.Project, int, error) {

	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.Project{
		project.AccountIDEQ(int(claims["account_id"].(float64))),
		project.DeletedEQ(false),
	}

	baseOrders := getProjectOrders(filters.Order)

	if queryParams.Query != "" {
		searchFilter := project.Or(
			project.NameContains(queryParams.Query),
			project.DescriptionContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	var wg sync.WaitGroup
	var projects []*generated.Project
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		projects, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(baseOrders...).
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

	return projects, totalFiltered, nil
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
func getProjectOrders(filters []models.OrderFields) []project.OrderOption {
	baseOrders := []project.OrderOption{}

	if len(filters) == 0 {
		baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderDesc()))
	}

	for _, sortCriteria := range filters {
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
		case "status":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByStatus(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByStatus(sql.OrderDesc()))
			}
		case "location":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByLocation(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByLocation(sql.OrderDesc()))
			}
		case "description":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByDescription(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByDescription(sql.OrderDesc()))
			}
		case "notes":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByNotes(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByNotes(sql.OrderDesc()))
			}
		case "start_date":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByStartDate(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByStartDate(sql.OrderDesc()))
			}
		case "end_date":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByEndDate(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByEndDate(sql.OrderDesc()))
			}
		case "deleted":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, project.ByDeleted(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, project.ByDeleted(sql.OrderDesc()))
			}
		default:
			// baseOrders = append(baseOrders, project.ByCreatedAt(sql.OrderDesc()))
			log.Println("Unknown Field is passed as order, ignore it")
			// Handle unknown fields or provide a default behavior
			// For example, you can log an error or ignore unknown fields
		}
	}

	return baseOrders
}
