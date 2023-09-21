package projects

import (
	"context"

	"github.com/Pyakz/buildbox-api/ent"
	"github.com/Pyakz/buildbox-api/ent/project"
	_ "github.com/lib/pq"
)

// ProjectService defines methods for project management.
type ProjectService interface {
	CreateProject(ctx context.Context, newProject *Project) (*ent.Project, error)
	GetProjects(ctx context.Context) ([]*ent.Project, error)
}

// EntProjectService is an implementation of ProjectService that uses the Ent client.
type projectService struct {
	client *ent.ProjectClient
}

// NewEntProjectService creates a new instance of EntProjectService.
func NewProjectService(client *ent.ProjectClient) ProjectService {
	return &projectService{client: client}
}

// CreateProject creates a new project.
func (s *projectService) CreateProject(ctx context.Context, newProject *Project) (*ent.Project, error) {

	// if status is not in the payload just use the default which is "planning"

	if newProject.Status == "" {
		newProject.Status = project.StatusPlanning
	}

	project, err := s.client.Create().
		// accountID can be found in claims
		// SetAccountID(newProject.AccountID).

		// createdBy can be found in claims
		// which is the id of the current user
		// SetCreatedBy(newProject.CreatedBy).

		SetName(newProject.Name).
		SetStatus(newProject.Status).
		SetStartDate(*newProject.StartDate).
		SetEndDate(*newProject.EndDate).
		SetClientID(newProject.ClientID).
		SetBudget(newProject.Budget).
		SetLocation(newProject.Location).
		SetDescription(newProject.Description).
		SetNotes(newProject.Notes).
		SetManagerID(newProject.ManagerID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *projectService) GetProjects(ctx context.Context) ([]*ent.Project, error) {
	projects, err := s.client.Query().Select(
		// return all fields
		project.FieldID,
		project.FieldAccountID,
		project.FieldClientID,
		project.FieldManagerID,
		project.FieldCreatedBy,
		project.FieldName,
		project.FieldDescription,
		project.FieldNotes,
		project.FieldStatus,
		project.FieldLocation,
		project.FieldBudget,
		project.FieldDeleted,
		project.FieldStartDate,
		project.FieldEndDate,
		project.FieldUpdatedAt,
		project.FieldCreatedAt,
	).All(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
