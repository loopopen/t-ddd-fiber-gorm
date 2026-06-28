package service

import (
	"context"
	"log/slog"

	"github.com/loopopen/t-ddd-fiber-gorm/internal/applic/query"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/applic/result"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/domain/event"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/domain/repo"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/pkg/x"

	"github.com/lopolopen/gap"
	"gorm.io/gorm"
)

type UserSvc struct {
	logger   *slog.Logger
	db       *gorm.DB
	pub      gap.EventPublisher
	userRepo repo.UserRepo
}

func NewUserSvc(logger *slog.Logger, db *gorm.DB, pub gap.EventPublisher, quantRepo repo.UserRepo) *UserSvc {
	svc := &UserSvc{
		logger:   x.SLogWithin(logger, &UserSvc{}),
		db:       db,
		pub:      pub,
		userRepo: quantRepo,
	}
	return svc
}

func (svc *UserSvc) Query(ctx context.Context, q query.UserQuery) ([]result.User, error) {
	users, err := svc.userRepo.Query(ctx, q.Key)
	if err != nil {
		return nil, err
	}
	var rs []result.User
	for _, q := range users {
		rs = append(rs, *new(result.User).FromEntity(q))
	}
	return rs, nil
}

//go:generate go tool gapc -file=$GOFILE

// @subscribe
func (svc *UserSvc) HandlerUserCreated() gap.Handler[*event.UserCreated] {
	return func(ctx context.Context, msg *event.UserCreated, headers map[string]string) error {
		panic("unimplemented")
	}
}
