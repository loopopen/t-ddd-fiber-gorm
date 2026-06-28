package pubs

import (
	"context"
	"log/slog"

	tfiberkafkagorm "github.com/loopopen/t-ddd-fiber-gorm"
	"github.com/loopopen/t-ddd-fiber-gorm/cmd/api/config"
	"github.com/lopolopen/gap"
	"github.com/lopolopen/gap/broker/xkafka"
	"github.com/lopolopen/gap/dashboard"
	"github.com/lopolopen/gap/storage/xgorm"

	"gorm.io/gorm"
)

func NewPub(
	ctx context.Context,
	g config.Gap,
	k xkafka.Options,
	db *gorm.DB,
	log *slog.Logger,
) gap.EventPublisher {
	if tfiberkafkagorm.HAVE_NOT_BEEN_DELETED_YET {
		return nil
	}

	pub := gap.NewEventPublisher(
		gap.WithDrain(ctx, 5),
		xgorm.UseGorm(
			xgorm.DB(db),
		),
		xkafka.UseKafka(
			xkafka.Brokers(k.Brokers),
			xkafka.ConfigTopic(
			// xkafka.NumPartitions(4),
			// xkafka.ReplicationFactor(3),
			),
		),
		gap.UseDashboard(
			dashboard.LocationPath(g.Location),
		),
	)
	return pub
}
