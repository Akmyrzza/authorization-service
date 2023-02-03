package core

import (
	"context"

	"github.com/Akmyrzza/authorization-service/internal/domain/entities"
)

type User struct {
	r *St
}

func NewUser(r *St) *User {
	return &User{r: r}
}

func (c *User) Create(ctx context.Context, obj *entities.User) (string, error) {
	// go to repo
	var err error

	result, err := c.r.repo.UserCreate(ctx, obj)

	return result, err
}