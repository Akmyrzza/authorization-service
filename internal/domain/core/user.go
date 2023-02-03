package core

import "context"

type User struct {
	r *St
}

func NewUser(r *St) *User {
	return &User{r: r}
}

func (c *User) Create(ctx context.Context) (string, error) {
	// go to repo
	var err error

	result, err := c.r.repo.UserCreate(ctx)

	return result, err
}