package persistence

import (
	"backend_crudgo/domain/users/domain/model"
	repoDomain "backend_crudgo/domain/users/domain/repository"
	"backend_crudgo/infrastructure/database"
	response "backend_crudgo/types"

	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type sqlUserRepo struct {
	Conn *database.DataDB
}

// NewUserRepository Should initialize the dependencies for this service.
func NewUserRepository(Conn *database.DataDB) repoDomain.UserRepository {
	return &sqlUserRepo{
		Conn: Conn,
	}
}

func (sr *sqlUserRepo) CreateUserHandler(ctx context.Context, user *model.User) (*response.CreateResponse, error) {
	var idResult string

	stmt, err := sr.Conn.DB.PrepareContext(ctx, InsertUser)
	if err != nil {
		return &response.CreateResponse{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()

	row := stmt.QueryRowContext(ctx, &user.UserID, &user.Name, &user.Email, &user.DateCreated)
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.CreateResponse{}, err
	}

	GenericUserResponse := response.CreateResponse{
		Message: "User created",
	}

	return &GenericUserResponse, nil
}

func (sr *sqlUserRepo) GetUserHandler(ctx context.Context, id string) (*response.GenericUserResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectUser)
	if err != nil {
		return &response.GenericUserResponse{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()

	row := stmt.QueryRowContext(ctx, id)
	user := &model.User{}

	err = row.Scan(&user.UserID, &user.Name, &user.Email, &user.DateCreated,
		&user.UserModify, &user.DateModify)
	if err != nil {
		return &response.GenericUserResponse{Error: err.Error()}, err
	}

	GenericUserResponse := &response.GenericUserResponse{
		Message: "Get user success",
		User:    user,
	}

	return GenericUserResponse, nil
}

func (sr *sqlUserRepo) GetUsersHandler(ctx context.Context) (*response.GenericUserResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectUsers)
	if err != nil {
		return &response.GenericUserResponse{}, nil
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()
	row, err := sr.Conn.DB.QueryContext(ctx, SelectUsers)

	var users []*model.User
	for row.Next() {
		var user = &model.User{}
		err = row.Scan(&user.UserID, &user.Name, &user.Email, &user.DateCreated,
			&user.UserModify, &user.DateModify)

		users = append(users, user)
	}
	if err != nil {
		return &response.GenericUserResponse{Error: err.Error()}, err
	}
	GenericUserResponse := &response.GenericUserResponse{
		Message: "Get user success",
		User:    users,
	}

	return GenericUserResponse, nil
}
