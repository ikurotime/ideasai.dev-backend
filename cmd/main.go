package main

import (
	. "ikurotime/ideasai/internal/infrastructure/http"
	. "ikurotime/ideasai/internal/infrastructure/http/route"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`)),
			AsRoute(NewHelloHandler),
			zap.NewExample),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
