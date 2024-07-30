package main

import (
	"ikurotime/ideasai/internal/application/handler"
	. "ikurotime/ideasai/internal/infrastructure/http"
	route "ikurotime/ideasai/internal/infrastructure/http/route"
	"ikurotime/ideasai/internal/infrastructure/repository"
	pkg "ikurotime/ideasai/pkg/query"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewHTTPRouterGinGonic,
				fx.ParamTags(`group:"routes"`)),
			AsRoute(route.NewFindProjectByIDRoute),
			fx.Annotate(
				pkg.NewInternalQueryBus,
				fx.ParamTags(`group:"queryHandler"`),
			),
			AsQueryHandler(handler.NewFindProjectHandler),
			zap.NewExample,
			repository.NewProjectRepository,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
