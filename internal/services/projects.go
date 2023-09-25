package services

import (
	"context"
	"errors"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/golang-jwt/jwt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// ProjectService defines methods for project management.
type ProjectService interface {
	CreateProject(ctx context.Context, newProject *generated.Project) (*generated.Project, error)
	GetProjects(ctx context.Context) ([]*generated.Project, error)
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
		// temporary
		SetClientID("random-client-id").
		SetName(newProject.Name).
		SetStartDate(newProject.StartDate).
		SetNillableStatus(newProject.Status).
		SetNillableEndDate(newProject.EndDate).
		SetBudget(newProject.Budget).
		SetNillableLocation(newProject.Location).
		SetNillableDescription(newProject.Description).
		SetNillableNotes(newProject.Notes).
		SetNillableManagerID(newProject.ManagerID).
		SetNillableDeleted(newProject.Deleted).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *projectService) GetProjects(ctx context.Context) ([]*generated.Project, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		// User claims not found in context, handle error accordingly
		return nil, errors.New("failed to get user claims from context")
	}
	projects, err := s.client.Query().Where(
		project.AccountIDEQ(int(claims["account_id"].(float64))),
	).All(ctx)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

// UpdateProject updates an existing project.
func (s *projectService) UpdateProjectByID(ctx context.Context, id int, newPayload *generated.Project) (*generated.Project, error) {
	project, err := s.client.UpdateOneID(id).
		SetName(newPayload.Name).
		SetStartDate(newPayload.StartDate).
		SetNillableStatus(newPayload.Status).
		SetNillableEndDate(newPayload.EndDate).
		SetBudget(newPayload.Budget).
		SetNillableLocation(newPayload.Location).
		SetNillableDescription(newPayload.Description).
		SetNillableNotes(newPayload.Notes).
		SetNillableManagerID(newPayload.ManagerID).
		SetNillableDeleted(newPayload.Deleted).
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
