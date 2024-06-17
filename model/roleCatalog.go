package model

type RoleCatalog struct {
	RoleId   string `json:"roleId" binding:"required"`
	RoleCode string `json:"roleCode" binding:"required"`
	RoleDesc string `json:"roleDesc"`
	Lob      string `json:"lob" binding:"required"`
}

func PopulateRoleCatalog() RoleCatalog {

	return RoleCatalog{RoleId: "abc123", RoleCode: "PR", RoleDesc: "Primary"}
}
