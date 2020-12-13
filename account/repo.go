package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

//se utiliza un mensaje de error que se va a mostrar al dar un error XD
var RepoErr = errors.New("Unablre to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

//el ejemplo esta con una base postgres
func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES( $1, $2, $3`

	//revisamos que exista un email y pass en el usuario y si no retornar error
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	//si todo esta bien insertamos en la base
	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}
	return email, nil
}
