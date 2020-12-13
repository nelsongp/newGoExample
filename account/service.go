package account

import (
	"context"
)

//La interfaz sirve para exponerle al transport la logica del negocio
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
