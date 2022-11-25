// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogs is the golang structure for table sys_logs.
type SysLogs struct {
	Id        int64       `json:"id"        description:"ID"`
	UserId    int64       `json:"userId"    description:"用户UID"`
	Error     string      `json:"error"     description:"错误信息"`
	Category  string      `json:"category"  description:"分类"`
	Level     int         `json:"level"     description:"等级"`
	Content   string      `json:"content"   description:"日志内容"`
	Context   string      `json:"context"   description:"上下文数据"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
