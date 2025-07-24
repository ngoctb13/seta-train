package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/handler/gql/user/directive"
	usergqlgenerated "github.com/ngoctb13/seta-train/handler/gql/user/generated"
	usergqlresolver "github.com/ngoctb13/seta-train/handler/gql/user/resolver"
	hdl "github.com/ngoctb13/seta-train/handler/rest"
	"github.com/ngoctb13/seta-train/infra/repos"
	"github.com/ngoctb13/seta-train/internal/domains/user/usecases"
)

type domains struct {
	user *usecases.User
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
	user := usecases.NewUser(repo.Users())
	return &domains{
		user: user,
	}
}

func (s *Server) initGqlRoute(domains *domains) {
	gqlHandler := handler.NewDefaultServer(usergqlgenerated.NewExecutableSchema(usergqlgenerated.Config{
		Resolvers: &usergqlresolver.Resolver{
			UserUsecase: domains.user,
		},
		Directives: usergqlgenerated.DirectiveRoot{
			Auth: directive.AuthDirective,
		},
	}))
	s.router.POST("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})
	s.router.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	})
}

func (s *Server) initRestRoute(domains *domains) {
	handler := hdl.NewHandler(domains.user)

	routerAuth := s.router.Group("v1")
	handler.ConfigAuthRouteAPI(routerAuth)
}

func (s *Server) initRouter(domains *domains, repo repos.IRepo) {
	//init gql route
	s.initGqlRoute(domains)

	//init rest route
	s.initRestRoute(domains)
}
