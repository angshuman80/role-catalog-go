package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"role-catalog-go/Controller"
)

type Routes struct {
	router *gin.Engine
}

func InitializeRoleCatalogRoute() Routes {
	r := Routes{router: gin.Default()}
	Controller.RolecatalogRoutes(r.router.Group("/v1"))
	return r
}

func (r *Routes) Run() {
	log.Fatalln(r.router.Run(":8090"))
}
