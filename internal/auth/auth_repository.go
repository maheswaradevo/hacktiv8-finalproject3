package auth

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type User interface {
	Save(ctx context.Context, data model.User) (uint64, error)
	CheckEmail(ctx context.Context, email string) (*model.User, error)
	UpdateAccount(ctx context.Context, data model.User, userID uint64) error
}
