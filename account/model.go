package account

import "context"

//se representa el objeto struct a ser utilizado en el transpor y service
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//esto es del repo, es como la inicializacion del repo
type Repository interface{
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
