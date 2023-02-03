package core

import "github.com/Akmyrzza/authorization-service/internal/adapters/repo"

type St struct {
	repo repo.Repo
	User *User
}