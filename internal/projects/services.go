package projects

import (
	"context"

	"github.com/Pyakz/buildbox-api/ent"
	"github.com/Pyakz/buildbox-api/ent/project"
	"github.com/Pyakz/buildbox-api/utils"
	_ "github.com/lib/pq"
)

// ProjectService defines methods for project management.
type ProjectService interface {
	CreateProject(ctx context.Context, newProject *ent.Project) (*ent.Project, error)
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
// CreateProject creates a new project.
func (s *projectService) CreateProject(ctx context.Context, newProject *ent.Project) (*ent.Project, error) {
	project, err := s.client.Create().
		// accountID and createdBy can be found in claims
		SetAccountID("id").
		SetCreatedBy("id").

		// temporary
		SetClientID("4").
		SetName(newProject.Name).
		SetStartDate(newProject.StartDate).
		SetStatus(newProject.Status).
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

func (s *projectService) GetProjects(ctx context.Context) ([]*ent.Project, error) {
	projects, err := s.client.Query().Select(
		// return all fields
		project.FieldID,
		project.FieldUUID,
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

	utils.ConsoleLog(projects)
	return projects, nil
}
