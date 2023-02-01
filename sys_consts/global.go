package sys_consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type global struct {
	UserDefaultType          sys_enum.UserType
	UserDefaultState         sys_enum.UserState
	NotAllowLoginUserTypeArr *garray.SortedIntArray
	LogLevelToDatabaseArr    *garray.SortedIntArray
	ApiPreFix                string
	OrmCacheConf             []*sys_model.TableCacheConf
	PermissionTree           []*permission.SysPermissionTree // PermissionTree 权限信息定义
	Searcher                 *xdb.Searcher
}

var (
	Global = global{
		UserDefaultType:          sys_enum.User.Type.SuperAdmin,
		UserDefaultState:         sys_enum.User.State.Normal,
		NotAllowLoginUserTypeArr: garray.NewSortedIntArray(),
		LogLevelToDatabaseArr:    garray.NewSortedIntArray(),
		ApiPreFix:                "",
		OrmCacheConf:             []*sys_model.TableCacheConf{},
		PermissionTree:           []*permission.SysPermissionTree{},
	}
)
