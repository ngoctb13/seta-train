package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/auth-service/graph"
	handlers "github.com/ngoctb13/seta-train/auth-service/handler"
	"github.com/ngoctb13/seta-train/auth-service/internal/auth"
	user_usecases "github.com/ngoctb13/seta-train/auth-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/auth-service/repos"
	"github.com/vektah/gqlparser/v2/ast"
)

type domains struct {
	user          *user_usecases.User
	importUsecase *user_usecases.ImportUsecase
}

func (s *Server) initCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) initDomains(repo repos.IRepo) *domains {
	user := user_usecases.NewUser(repo.Users())
	importUsecase := user_usecases.NewImportUsecase(repo.Users())
	return &domains{
		user:          user,
		importUsecase: importUsecase,
	}
}

func (s *Server) initAuthRoute(domains *domains) {
	gqlServer := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserUsecase: domains.user,
			Logger:      s.logger,
		},
		Directives: graph.DirectiveRoot{
			Auth: graph.AuthDirective,
		},
	}))

	gqlServer.AddTransport(transport.Options{})
	gqlServer.AddTransport(transport.GET{})
	gqlServer.AddTransport(transport.POST{})

	gqlServer.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	gqlServer.Use(extension.Introspection{})
	gqlServer.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	importHandler := handlers.NewImportHandler(domains.importUsecase, s.logger)

	s.router.POST("/query", func(c *gin.Context) {
		auth.AuthContextMiddleware(gqlServer).ServeHTTP(c.Writer, c.Request)
	})
	s.router.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))

	s.router.POST("/import-users", func(c *gin.Context) {
		importHandler.ImportUsers(c.Writer, c.Request)
	})
}
