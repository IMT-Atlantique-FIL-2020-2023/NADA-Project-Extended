package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/config"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/generated"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	"github.com/henvic/httpretty"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	herokuPort := os.Getenv("PORT")
	port := fmt.Sprint(config.CurrentConfig.Port)
	if herokuPort != "" {
		port = herokuPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(), Directives: generated.DirectiveRoot{
		Examples: func(ctx context.Context, obj interface{}, next graphql.Resolver, values []*string) (res interface{}, err error) {
			return next(ctx)
		},
		Fake: func(ctx context.Context, obj interface{}, next graphql.Resolver, typeArg model.FakeTypes, options *model.FakeOptions, locale *model.FakeLocale) (res interface{}, err error) {
			return next(ctx)
		},
		ListLength: func(ctx context.Context, obj interface{}, next graphql.Resolver, min, max int) (res interface{}, err error) {
			return next(ctx)
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger := &httpretty.Logger{
		Time:           true,
		TLS:            true,
		RequestHeader:  true,
		RequestBody:    true,
		ResponseHeader: true,
		ResponseBody:   true,
		Colors:         true, // erase line if you don't like colors
		Formatters:     []httpretty.Formatter{&httpretty.JSONFormatter{}},
	}
	logger.Middleware(srv)

	log.Info().Msgf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err((http.ListenAndServe(":"+port, nil)))
}
