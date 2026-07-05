//go:build wireinject

package main

import (
	"context"
	"log/slog"

	"github.com/loopopen/t-ddd-fiber-gorm/cmd/api/config"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/adapters/http"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/adapters/outbound"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/applic/service"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/infra/conf"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/infra/gorm"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/infra/gorm/repoimpl"
	"github.com/lopolopen/gap/broker/xkafka"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func wireApp(ctx context.Context, e *config.Env, c *config.Config, g config.Gap, k xkafka.Options, orm conf.ORM, log *slog.Logger) (*fiber.App, error) {
	panic(wire.Build(
		repoimpl.ProviderSet,
		service.ProviderSet,
		outbound.ProviderSet,
		gorm.NewGormDB,
		http.ProviderSet,
	))
}
