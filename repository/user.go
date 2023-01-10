package repository

import (
	"context"

	"github.com/cecardev/go-rest-server/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) (id int64,err error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
    GetUserByEmail(ctx context.Context, email string) (*models.User, error)
    Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (id int64,err error) {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
    return implementation.Close()
}

func GetUserByEmail(ctx context.Context, email string) (*models.User,error){
    return implementation.GetUserByEmail(ctx, email)
}
