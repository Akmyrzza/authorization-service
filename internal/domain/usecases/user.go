package usecases

import "context"

func (u *St) UserCreate(ctx context.Context) (string, error) {
	//go to core
	var err error
	var result string

	//err = u.db.
	result, err = u.cr.User.Create(ctx)
	
	return result, err
}