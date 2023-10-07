package services

import (
	"context"
	"errors"
	"log"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

type UserService interface {
	RegisterUser(ctx context.Context, newUser *generated.User) (*generated.User, error)
	GetUserByEmail(ctx context.Context, email string) (*generated.User, error)
	GetUserByID(ctx context.Context, id int) (*generated.User, error)
	GetUsers(ctx context.Context, queryParams *render.QueryParams, filters models.Filters) ([]*generated.User, int, error)
}

type userService struct {
	client *generated.UserClient
}

func NewUserService(client *generated.UserClient) UserService {
	return &userService{client: client}
}

func (s *userService) RegisterUser(ctx context.Context, newUser *generated.User) (*generated.User, error) {

	user, err := s.client.Create().
		SetAccountID(newUser.AccountID).
		SetEmail(newUser.Email).
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetPassword(newUser.Password).
		SetNillableMiddleName(newUser.MiddleName).
		SetNillableBirthday(newUser.Birthday).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*generated.User, error) {
	user, err := s.client.Query().Where(
		user.EmailEQ(email),
	).WithAccount().First(ctx)

	if err != nil {
		if generated.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("something went wrong, please try again later")
	}

	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*generated.User, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get user claims from context")
	}

	user, err := s.client.
		Query().
		Where(
			user.AccountIDEQ(int(claims["account_id"].(float64))),
			user.IDEQ(id),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(ctx context.Context, queryParams *render.QueryParams, filters models.Filters) ([]*generated.User, int, error) {
	claims, ok := ctx.Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		return nil, 0, errors.New("failed to get user claims from context")
	}

	baseFilters := []predicate.User{
		user.AccountIDEQ(int(claims["account_id"].(float64))),
	}

	if queryParams.Query != "" {
		searchFilter := user.Or(
			user.FirstNameContains(queryParams.Query),
			user.MiddleNameContains(queryParams.Query),
			user.LastNameContains(queryParams.Query),
			user.EmailContains(queryParams.Query),
		)
		baseFilters = append(baseFilters, searchFilter)
	}

	baseOrders := getUserOrder(filters.Order)

	var wg sync.WaitGroup
	var roles []*generated.User
	var totalFiltered int
	var err1, err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		roles, err1 = s.client.
			Query().
			Where(baseFilters...).
			Order(
				baseOrders...,
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

	return roles, totalFiltered, nil
}

func getUserOrder(filters []models.OrderFields) []user.OrderOption {
	baseOrders := []user.OrderOption{}

	if len(filters) == 0 {
		baseOrders = append(baseOrders, user.ByCreatedAt(sql.OrderDesc()))
	}

	for _, sortCriteria := range filters {
		switch sortCriteria.Field {
		case "first_name":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByFirstName(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByFirstName(sql.OrderDesc()))
			}
		case "middle_name":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByMiddleName(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByMiddleName(sql.OrderDesc()))
			}
		case "last_name":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByLastName(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByLastName(sql.OrderDesc()))
			}
		case "email":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByEmail(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByEmail(sql.OrderDesc()))
			}
		case "created_at":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByCreatedAt(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByCreatedAt(sql.OrderDesc()))
			}
		case "updated_at":
			if sortCriteria.Direction == "asc" {
				baseOrders = append(baseOrders, user.ByUpdatedAt(sql.OrderAsc()))
			} else if sortCriteria.Direction == "desc" {
				baseOrders = append(baseOrders, user.ByUpdatedAt(sql.OrderDesc()))
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
