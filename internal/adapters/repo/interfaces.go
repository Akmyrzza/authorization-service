package repo

import (
	"context"

	"github.com/Akmyrzza/authorization-service/internal/domain/entities"
)

type Repo interface {
	// user
	UserCreate(ctx context.Context, obj *entities.User) (string, error)
}