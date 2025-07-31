package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ngoctb13/seta-train/auth-service/graph"
	handlers "github.com/ngoctb13/seta-train/auth-service/handler"
	"github.com/ngoctb13/seta-train/auth-service/internal/auth"
	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/auth-service/repos"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/infra"
	"github.com/ngoctb13/seta-train/shared-modules/logger"
	"github.com/ngoctb13/seta-train/shared-modules/setting"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {

	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	// Initialize logger
	logger := logger.InitLogger("auth-service")

	cfg, err := config.Load(configFile)
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		panic(err)
	}

	// connect to db
	go setting.ConnectDatabase(cfg.DB)

	db, err := infra.InitPostgres(cfg.DB)
	if err != nil {
		logger.Error("Failed to initialize database: %v", err)
		panic(err)
	}

	repos := repos.NewSQLRepo(db, cfg.DB)
	userRepo := repos.Users()
	userUsecase := usecases.NewUser(userRepo)
	importUsecase := usecases.NewImportUsecase(userRepo)

	importHandler := handlers.NewImportHandler(importUsecase, logger)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserUsecase: userUsecase,
			Logger:      logger,
		},
		Directives: graph.DirectiveRoot{
			Auth: graph.AuthDirective,
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.AuthContextMiddleware(srv))
	http.HandleFunc("/import-users", importHandler.ImportUsers)

	logger.Info("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
