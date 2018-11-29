package oauth

import (
	"github.com/RichardKnop/go-oauth2-server/util/routes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

const (
	tokensResource     = "tokens"
	tokensPath         = "/" + tokensResource
	introspectResource = "introspect"
	introspectPath     = "/" + introspectResource
)

// RegisterRoutes registers route handlers for the oauth service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRoutes(s.GetRoutes(), subRouter)
}

// GetRoutes returns []routes.Route slice for the oauth service
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "oauth_tokens",
			Method:      "POST",
			Pattern:     tokensPath,
			HandlerFunc: s.tokensHandler,
		},
		{
			Name:        "oauth_introspect",
			Method:      "POST",
			Pattern:     introspectPath,
			HandlerFunc: s.introspectHandler,
		},
	}
}

func (s *Service) RegisterGinRoutes(router *gin.Engine, prefix string) {
	subRouter := router.Group(prefix)
	subRouter.POST(tokensPath, gin.WrapF(s.tokensHandler))
	subRouter.POST(introspectPath, gin.WrapF(s.introspectHandler))
}
