package outbound

import (
	"github.com/google/wire"
	"github.com/loopopen/t-ddd-fiber-gorm/internal/adapters/outbound/pubs"
)

var ProviderSet = wire.NewSet(pubs.NewPub)
