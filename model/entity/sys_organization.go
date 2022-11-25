// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysOrganization is the golang structure for table sys_organization.
type SysOrganization struct {
	Id          int64  `json:"id"          description:""`
	Name        string `json:"name"        description:"名称"`
	ParentId    int64  `json:"parentId"    description:"父级ID"`
	CascadeDeep int    `json:"cascadeDeep" description:"级联深度"`
	Description string `json:"description" description:"描述"`
}
