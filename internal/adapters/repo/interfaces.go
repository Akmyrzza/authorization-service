package repo

import "context"

type Repo interface {
	// user
	UserCreate(ctx context.Context) (string, error)
}