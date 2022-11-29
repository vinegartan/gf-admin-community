package daoctl

import (
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
)

func RemoveQueryCache(db gdb.DB, prefix string) {
	prefix = "SelectCache:" + prefix
	cacheKeys, _ := db.GetCache().KeyStrings(db.GetCtx())
	for _, key := range cacheKeys {
		if gstr.HasPrefix(key, prefix) {
			db.GetCache().Remove(db.GetCtx(), key)
		}
	}
}

func GetById[T any](db *gdb.Model, id int64) *T {
	result := new(T)

	if err := db.Where("id", id).Scan(result); err != nil {
		return nil
	}
	return result
}

func Count(db *gdb.Model, orderBy *model.OrderBy, searchFields ...model.SearchField) (total int) {
	db, err := makeBuilder(db, orderBy, searchFields)
	if err != nil {
		return 0
	}
	total, _ = db.Count()
	return
}

func builder(db *gdb.Model, orderBy *model.OrderBy, searchFields ...model.SearchField) (*gdb.Model, error) {
	return makeBuilder(db, orderBy, searchFields)
}

func makeBuilder(db *gdb.Model, orderBy *model.OrderBy, searchFieldArr []model.SearchField) (*gdb.Model, error) {
	if len(searchFieldArr) <= 0 {
		return nil, gerror.New("查询条件参数错误")
	}
	for index, field := range searchFieldArr {
		field.Field = gstr.CaseSnakeFirstUpper(field.Field)
		if gconv.String(field.Field) != "" {
			// 过滤特殊符号，防止SQL注入
			orderBy.Fields = gstr.ReplaceIByMap(orderBy.Fields, map[string]string{
				"\"": "",
				"'":  "",
			})

			if index == 0 {
				field.IsOrWhere = false
			}

			if field.IsOrWhere {
				if gstr.CaseCamelLower(field.Where) == "in" {
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereOrNotIn(field.Field, field.Value)
					} else {
						db = db.WhereOrIn(field.Field, field.Value)
					}
				} else if gstr.CaseCamelLower(field.Where) == "between" {
					valueArr := gstr.SplitAndTrim(gconv.String(field.Value), ",")
					minValue := valueArr[0]
					maxValue := minValue
					if len(valueArr) > 1 {
						maxValue = valueArr[1]
					}
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereOrNotBetween(field.Field, minValue, maxValue)
					} else {
						db = db.WhereOrBetween(field.Field, minValue, maxValue)
					}
				} else if gstr.CaseCamelLower(field.Where) == "like" {
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereOrNotLike(field.Field, field.Value)
					} else {
						db = db.WhereOrLike(field.Field, field.Value)
					}
				} else {
					if field.Where == ">" {
						db = db.WhereOrGT(field.Field, field.Value)
					} else if field.Where == ">=" {
						db = db.WhereOrGTE(field.Field, field.Value)
					} else if field.Where == "<" {
						db = db.WhereOrLT(field.Field, field.Value)
					} else if field.Where == "<=" {
						db = db.WhereOrLT(field.Field, field.Value)
					} else if field.Where == "<>" {
						db = db.WhereOrNotIn(field.Field, field.Value)
					} else if field.Where == "=" {
						db = db.WhereOr(field.Field, field.Value)
					} else {
						return nil, gerror.New("查询条件参数错误")
					}
				}
			} else {
				if gstr.CaseCamelLower(field.Where) == "in" {
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereNotIn(field.Field, field.Value)
					} else {
						db = db.WhereIn(field.Field, field.Value)
					}
				} else if gstr.CaseCamelLower(field.Where) == "between" {
					valueArr := gstr.SplitAndTrim(gconv.String(field.Value), ",")
					minValue := valueArr[0]
					maxValue := minValue
					if len(valueArr) > 1 {
						maxValue = valueArr[1]
					}
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereNotBetween(field.Field, minValue, maxValue)
					} else {
						db = db.WhereBetween(field.Field, minValue, maxValue)
					}
				} else if gstr.CaseCamelLower(field.Where) == "like" {
					if gstr.CaseCamelLower(field.Modifier) == "not" {
						db = db.WhereNotLike(field.Field, field.Value)
					} else {
						db = db.WhereLike(field.Field, gconv.String(field.Value))
					}
				} else {
					if field.Where == ">" {
						db = db.WhereGT(field.Field, field.Value)
					} else if field.Where == ">=" {
						db = db.WhereGTE(field.Field, field.Value)
					} else if field.Where == "<" {
						db = db.WhereLT(field.Field, field.Value)
					} else if field.Where == "<=" {
						db = db.WhereLT(field.Field, field.Value)
					} else if field.Where == "<>" {
						db = db.WhereNotIn(field.Field, field.Value)
					} else if field.Where == "=" {
						db = db.Where(field.Field, field.Value)
					} else {
						return nil, gerror.New("查询条件参数错误")
					}
				}
			}
		}
	}
	if orderBy != nil && len(orderBy.Fields) > 0 {
		// 过滤特殊符号，防止SQL注入
		orderBy.Fields = gstr.ReplaceIByMap(orderBy.Fields, map[string]string{
			"\"": "",
			"'":  "",
		})

		if gstr.CaseCamelLower(orderBy.Sort) == "asc" || len(orderBy.Sort) <= 0 {
			db = db.OrderAsc(orderBy.Fields)
		} else {
			db = db.OrderDesc(orderBy.Fields)
		}
	}
	return db, nil
}

func Query[T any](db *gdb.Model, searchFields *model.SearchFilter, orderBy *model.OrderBy, IsExport bool) (response *model.CollectRes[T], err error) {
	itemsDb, err := makeBuilder(db, orderBy, searchFields.Fields)
	if err != nil {
		return nil, err
	}

	total, _ := itemsDb.Count()

	queryDb, _ := makeBuilder(db, orderBy, searchFields.Fields)

	entities := make([]T, 0)
	if searchFields == nil || IsExport {
		err = queryDb.Scan(&entities)
	} else {
		err = queryDb.Page(searchFields.Page, searchFields.PageSize).Scan(&entities)
	}

	response = &model.CollectRes[T]{
		List: &entities,
		PaginationRes: model.PaginationRes{
			Pagination: searchFields.Pagination,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(searchFields.PageSize))),
		},
	}

	return response, nil
}

func Find[T any](db *gdb.Model, orderBy *model.OrderBy, searchFields ...model.SearchField) (response *model.CollectRes[T], err error) {
	return Query[T](db, &model.SearchFilter{
		Fields: searchFields,
		Pagination: model.Pagination{
			Page:     1,
			PageSize: 1000,
		},
	}, orderBy, true)
}

func GetAll[T any](db *gdb.Model, info *model.Pagination) (response *model.CollectRes[T], err error) {
	total, err := db.Count()
	entities := make([]T, 0, total)
	if info == nil {
		info = &model.Pagination{
			Page:     1,
			PageSize: total,
		}
	}

	if err != nil {
		return
	}
	err = db.Page(info.Page, info.PageSize).Scan(&entities)

	return &model.CollectRes[T]{
		List: &entities,
		PaginationRes: model.PaginationRes{
			Pagination: *info,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(info.PageSize))),
		},
	}, nil
}
