// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FdInvoiceDao is the data access object for table fd_invoice.
type FdInvoiceDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FdInvoiceColumns // columns contains all the column names of Table for convenient usage.
}

// FdInvoiceColumns defines and stores column names for table fd_invoice.
type FdInvoiceColumns struct {
	Id             string //
	Name           string // 发票抬头名称
	TaxId          string // 纳税识别号
	Addr           string // 发票收件地址，限纸质
	Email          string // 发票收件邮箱，限电子发票
	UserId         string // 申请人UserID
	AuditUserId    string // 审核人UserID
	AuditReplayMsg string // 审核回复，仅审核不通过时才有值
	AuditAt        string // 审核时间
	State          string // 状态：1待审核、2已通过、3不通过
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// fdInvoiceColumns holds the columns for table fd_invoice.
var fdInvoiceColumns = FdInvoiceColumns{
	Id:             "id",
	Name:           "name",
	TaxId:          "tax_id",
	Addr:           "addr",
	Email:          "email",
	UserId:         "user_id",
	AuditUserId:    "audit_user_id",
	AuditReplayMsg: "audit_replay_msg",
	AuditAt:        "audit_at",
	State:          "state",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewFdInvoiceDao creates and returns a new DAO object for table data access.
func NewFdInvoiceDao() *FdInvoiceDao {
	return &FdInvoiceDao{
		group:   "default",
		table:   "fd_invoice",
		columns: fdInvoiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FdInvoiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FdInvoiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FdInvoiceDao) Columns() FdInvoiceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FdInvoiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FdInvoiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FdInvoiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
