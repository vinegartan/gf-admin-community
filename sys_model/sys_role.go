package sys_model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type SysRole struct {
	Id          int64  `p:"id"         dc:"ID"`
	Name        string `p:"name"       v:"required|length:1,16#请输入角色名称|角色名称长度限定1~16字符" dc:"角色名称"`
	IsSystem    bool   `json:"isSystem"    description:"是否默认角色，true仅能修改名称，不允许删除和修改"`
	Description string `p:"description"       v:"max-length:128#角色描述长度限定128字符" dc:"角色描述"`
	UnionMainId int64  `p:"unionMainId" v:"required|#请输入该角色所属主体" dc:"主体id"`
}

type RoleListRes CollectRes[sys_entity.SysRole]

type RoleHookFunc func(ctx context.Context, userInfo sys_entity.SysUser) (int64, error)

type RoleHookInfo HookEventType[sys_enum.UserType, RoleHookFunc]
