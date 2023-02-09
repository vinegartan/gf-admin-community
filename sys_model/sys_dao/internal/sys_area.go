// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAreaDao is the data access object for table sys_area.
type SysAreaDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysAreaColumns // columns contains all the column names of Table for convenient usage.
}

// SysAreaColumns defines and stores column names for table sys_area.
type SysAreaColumns struct {
	Id            string // ID
	AreaCode      string // 地区编码
	AreaName      string // 地区名称
	Level         string // 1:省份province,2:市city,3:区县district,4:街道street
	CityCode      string // 城市编码
	LatLongCenter string // 城市中心点（即经纬度）
	ParentId      string // 地区父节点
}

// sysAreaColumns holds the columns for table sys_area.
var sysAreaColumns = SysAreaColumns{
	Id:            "id",
	AreaCode:      "area_code",
	AreaName:      "area_name",
	Level:         "level",
	CityCode:      "city_code",
	LatLongCenter: "lat_long_center",
	ParentId:      "parent_id",
}

// NewSysAreaDao creates and returns a new DAO object for table data access.
func NewSysAreaDao(proxy ...dao_interface.IDao) *SysAreaDao {
	var dao *SysAreaDao
	if proxy != nil {
		dao = &SysAreaDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysAreaColumns,
		}
		return dao
	}

	return &SysAreaDao{
		group:   "default",
		table:   "sys_area",
		columns: sysAreaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAreaDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAreaDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysAreaDao) Columns() SysAreaColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAreaDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	model := dao.DB().Model(dao.Table()).Safe().Ctx(ctx)

	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		Model: model,
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
		}
	}

	model = daoctl.RegisterDaoHook(model)

	return model
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
