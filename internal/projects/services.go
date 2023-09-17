package projects

import (
	"context"
	"fmt"
	"time"

	"github.com/Pyakz/buildbox-api/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectService interface {
	CreateProject(ctx context.Context, newProject *Project) (*Project, error)
	GetProjectByID(ctx context.Context, projectID string) (*Project, error)

	// Add more methods here for project management
}

type projectService struct {
	collection *mongo.Collection
	// You can include other dependencies or repositories here if needed
}

func NewProjectService(collection *mongo.Collection) ProjectService {
	return &projectService{
		collection: collection,
	}
}

func (p *projectService) CreateProject(ctx context.Context, newProject *Project) (*Project, error) {
	if newProject.ID == "" {
		newProject.ID = utils.GenerateID()
	}

	if newProject.StartDate == nil {
		currentDate := time.Now().UTC()
		newProject.StartDate = &currentDate
	}

	if newProject.Status == "" {
		newProject.Status = "planning" // Default status
	}

	if _, err := p.collection.InsertOne(ctx, newProject); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return newProject, nil
}

func (p *projectService) GetProjectByID(ctx context.Context, projectID string) (*Project, error) {
	// Validate that projectID is a valid UUID
	_, err := uuid.Parse(projectID)
	if err != nil {
		return nil, fmt.Errorf("invalid project ID: %w", err)
	}

	project := &Project{}

	if err := p.collection.FindOne(ctx, bson.M{"_id": projectID}).Decode(project); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Project not found for the provided ID")
		}
		return nil, fmt.Errorf("an error occurred while retrieving the project: %w", err)
	}
	println("--", project)
	return project, nil
}
