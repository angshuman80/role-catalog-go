package main

import (
	"role-catalog-go/router"
)

func main() {
	routes := router.InitializeRoleCatalogRoute()
	routes.Run()
}
