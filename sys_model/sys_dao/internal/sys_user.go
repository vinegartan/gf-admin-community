// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for table sys_user.
type SysUserColumns struct {
	Id        string //
	Username  string // 账号
	Password  string // 密码
	State     string // 状态：0未激活、1正常、-1封号、-2异常、-3已注销
	Type      string // 用户类型，0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心，64后台
	Mobile    string // 手机号
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// sysUserColumns holds the columns for table sys_user.
var sysUserColumns = SysUserColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	State:     "state",
	Type:      "type",
	Mobile:    "mobile",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(proxy ...dao_interface.IDao) *SysUserDao {
	var dao *SysUserDao
	if proxy != nil {
		dao = &SysUserDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysUserColumns,
		}
		return dao
	}

	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysUserDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
