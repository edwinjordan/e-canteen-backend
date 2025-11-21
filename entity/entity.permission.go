package entity

type Permission struct {
	PermissionId          int    `json:"permission_id"`
	PermissionName        string `json:"permission_name"`
	PermissionResource    string `json:"permission_resource"`
	PermissionAction      string `json:"permission_action"`
	PermissionDescription string `json:"permission_description,omitempty"`
	PermissionStatus      string `json:"permission_status"`
	PermissionParentId    *int   `json:"permission_parent_id,omitempty"`
}

type PermissionResponse struct {
	PermissionId          int    `json:"permission_id"`
	PermissionName        string `json:"permission_name"`
	PermissionResource    string `json:"permission_resource"`
	PermissionAction      string `json:"permission_action"`
	PermissionDescription string `json:"permission_description,omitempty"`
	PermissionStatus      string `json:"permission_status"`
	PermissionParentId    *int   `json:"permission_parent_id,omitempty"`
}

type CreatePermissionRequest struct {
	PermissionName        string `json:"permission_name" validate:"required"`
	PermissionResource    string `json:"permission_resource" validate:"required"`
	PermissionAction      string `json:"permission_action" validate:"required"`
	PermissionDescription string `json:"permission_description,omitempty"`
	PermissionStatus      string `json:"permission_status" validate:"required,oneof=main_menu submenu action"`
	PermissionParentId    *int   `json:"permission_parent_id,omitempty"`
}

type UpdatePermissionRequest struct {
	PermissionName        string `json:"permission_name,omitempty"`
	PermissionResource    string `json:"permission_resource,omitempty"`
	PermissionAction      string `json:"permission_action,omitempty"`
	PermissionDescription string `json:"permission_description,omitempty"`
	PermissionStatus      string `json:"permission_status,omitempty" validate:"omitempty,oneof=main_menu submenu action"`
	PermissionParentId    *int   `json:"permission_parent_id,omitempty"`
}

// PermissionTreeNode represents a hierarchical permission structure
type PermissionTreeNode struct {
	PermissionId          int                    `json:"permission_id"`
	PermissionName        string                 `json:"permission_name"`
	PermissionResource    string                 `json:"permission_resource"`
	PermissionAction      string                 `json:"permission_action"`
	PermissionDescription string                 `json:"permission_description,omitempty"`
	PermissionStatus      string                 `json:"permission_status"`
	PermissionParentId    *int                   `json:"permission_parent_id,omitempty"`
	Children              []*PermissionTreeNode  `json:"children,omitempty"`
}
