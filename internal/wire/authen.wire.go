//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitAuthenRouterHandler() (*controller.AuthenController, error) {
	wire.Build(
		repo.NewAuthenRepo,
		service.NewAuthenService,
		controller.NewAuthenController,
	)

	return new(controller.AuthenController), nil
}
