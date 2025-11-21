package entity

type PermissionRole struct {
	PermissionId int `json:"permission_id"`
	RoleId       int `json:"role_id"`
}

type AssignPermissionRequest struct {
	RoleId        int   `json:"role_id" validate:"required"`
	PermissionIds []int `json:"permission_ids" validate:"required,min=1"`
}

type RevokePermissionRequest struct {
	RoleId       int `json:"role_id" validate:"required"`
	PermissionId int `json:"permission_id" validate:"required"`
}
