package persistence

import (
	"backend_crudgo/domain/users/domain/model"
	repoDomain "backend_crudgo/domain/users/domain/repository"
	"backend_crudgo/infrastructure/database"
	response "backend_crudgo/types"
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
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
	user.UserPassword = hashPassword(user.UserPassword)
	row := stmt.QueryRowContext(ctx, &user.UserID, &user.Name, &user.UserIdentifier, &user.Email,
		&user.UserPassword, &user.UserTypeIdentifier)
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.CreateResponse{}, err
	}

	GenericUserResponse := response.CreateResponse{
		Message: "User created",
	}

	return &GenericUserResponse, nil
}

func (sr *sqlUserRepo) LoginUserHandler(ctx context.Context, user *model.User) (*response.GenericUserResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectLoginUser)
	if err != nil {
		return &response.GenericUserResponse{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()

	row := stmt.QueryRowContext(ctx, user.Name)
	currentUser := &model.User{}

	err = row.Scan(&currentUser.UserID, &currentUser.Name, &currentUser.Email, &currentUser.UserIdentifier,
		&currentUser.UserPassword, &currentUser.UserTypeIdentifier)

	if err != nil {
		return &response.GenericUserResponse{Error: err.Error()}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.UserPassword), []byte(user.UserPassword))

	if err != nil {
		return &response.GenericUserResponse{Error: "password incorrect"}, nil
	}
	token, nil := generateToken(currentUser.UserID)
	// La contraseña es válida, devolver el usuario autenticado
	GenericUserResponse := &response.GenericUserResponse{
		Message: "success",
		User:    token,
	}

	return GenericUserResponse, nil
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

	err = row.Scan(&user.UserID, &user.Name, &user.Email, &user.UserIdentifier, &user.UserPassword, &user.DateCreated,
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
		err = row.Scan(&user.UserID, &user.Name, &user.Email, &user.UserIdentifier, &user.UserPassword,
			&user.DateCreated, &user.UserModify, &user.DateModify)

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

func hashPassword(password string) string {
	// Generate a hash of the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func generateToken(userID string) (string, error) {
	// Crear una estructura de claims para el token
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Firma el token con la clave secreta y devuelve el resultado
	secretKey := os.Getenv("SECRET_KEY")
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

/*
func updateUserPassword(db *sql.DB, user *User, newPassword string) error {
    // Generar hash a partir de la nueva contraseña usando bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Actualizar la contraseña en la base de datos
    _, err = db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, user.ID)
    if err != nil {
        return err
    }

    // Actualizar la contraseña en la referencia del objeto User
    user.Password = string(hashedPassword)

    return nil
}
*/

/* func authenticateUser(db *sql.DB, username string, password string) (*User, error) {
	// Buscar el usuario en la base de datos por nombre de usuario
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	// Comparar la contraseña en texto plano con el hash de contraseña cifrado almacenado en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// La contraseña es válida, devolver el usuario autenticado
	return &user, nil
} */
