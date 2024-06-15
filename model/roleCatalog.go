package model

type roleCatalog struct {
   RoleId string  `json:"roleId"`
   RoleCode string `json:"roleCode"`
   RoleDesc string  `json:"roleDesc"`
}

func populateRoleCatalog() roleCatalog{
  return  roleCatalog{RoleId: "abc123",RoleCode: "PR",RoleDesc: "Primary"}
}