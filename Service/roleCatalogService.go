package Service

import (
	"github.com/google/uuid"
	"log"
	"role-catalog-go/inits"
	"role-catalog-go/model"
)

func CreateRoleCatalog(catalog model.RoleCatalog) {
	roleId := uuid.New().String()
	roleCatalogTable := model.RoleCatalogTable{RoleId: roleId, RoleCode: catalog.RoleCode,
		RoleDesc: catalog.RoleDesc, Lob: catalog.Lob}
	roleByLobTable := model.RoleByLobTable{Lob: catalog.Lob, RoleIds: []string{roleId}}
	db := inits.GetDBInstance()
	db.Create(&roleCatalogTable)
	db.Create(&roleByLobTable)
	db.Commit()
}

func GetAllRoles() []model.RoleCatalog {
	db := inits.GetDBInstance()
	var catalog []model.RoleCatalogTable
	results := db.Find(&catalog)
	log.Println("results --> ", results)
	return nil
}
