// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysMenuDao is internal type for wrapping internal DAO implements.
type internalSysMenuDao = *internal.SysMenuDao

// sysMenuDao is the data access object for table sys_menu.
// You can define custom methods on it to extend its functionality as you wish.
type sysMenuDao struct {
	internalSysMenuDao
}

var (
	// SysMenu is globally public accessible object for table sys_menu operations.
	SysMenu = sysMenuDao{
		internal.NewSysMenuDao(),
	}
)

// Fill with you ideas below.
