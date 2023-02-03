package usecases

import (
	"github.com/Akmyrzza/authorization-service/internal/domain/core"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg logger.WarnAndError
	db string
	cr *core.St
}