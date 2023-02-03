package usecases

import (
	"context"

	"github.com/Akmyrzza/authorization-service/internal/domain/entities"
)

func (u *St) UserCreate(ctx context.Context, obj *entities.User) (string, error) {
	//go to core
	var err error
	var result string

	//err = u.db.
	result, err = u.cr.User.Create(ctx, obj)
	
	return result, err
}