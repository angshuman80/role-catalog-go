package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"role-catalog-go/Service"
	"role-catalog-go/model"
)

var rolecatalog model.RoleCatalog

func RolecatalogRoutes(rg *gin.RouterGroup) {
	rCatalog := rg.Group("/rolecatalog")
	rCatalog.GET("/", getAllRoleCatalog)
	rCatalog.POST("/", createRoleCatalog)
	rCatalog.GET("/:id", rolecatalogById)
}

func rolecatalogById(c *gin.Context) {
	log.Println("Parameter --> " + c.Param("id") + " role catalog id " + rolecatalog.RoleId)
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}
	if rolecatalog.RoleId == id {
		c.JSON(http.StatusOK, gin.H{"data": rolecatalog})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "role catalog not found"})
	}
}

func createRoleCatalog(c *gin.Context) {
	err := c.BindJSON(&rolecatalog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Println(rolecatalog.RoleId + " " + rolecatalog.RoleCode + " " + rolecatalog.Lob)
	Service.CreateRoleCatalog(rolecatalog)
	c.JSON(http.StatusCreated, "created")
}

func getAllRoleCatalog(c *gin.Context) {
	Service.GetAllRoles()
	c.JSON(200, gin.H{"data": rolecatalog})
}
