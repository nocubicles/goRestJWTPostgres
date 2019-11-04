package repository

import (
	"context"
	"goRestJWTPostgres/models"
)

type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	Create(ctx context.Context, u *models.User) (int64, error)
	Update(ctx context.Context, u *models.User) (*models.User, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type AuthRepo interface {
	Create(email string, password string) (string, error)
}
