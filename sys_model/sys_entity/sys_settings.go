// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSettings is the golang structure for table sys_settings.
type SysSettings struct {
	Name        string      `json:"name"        description:"配置名称"`
	Values      string      `json:"values"      description:"配置信息JSON格式"`
	Desc        string      `json:"desc"        description:"描述"`
	UnionMainId int64       `json:"unionMainId" description:"关联的主体id，为0代表是平台配置"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
}
