package funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func CheckPermission[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*sys_model.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermission(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
}
func CheckPermissionOr[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*sys_model.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOr(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
}

func CheckPersonLicenseFiles[T sys_entity.SysPersonLicense | sys_do.SysPersonLicense](ctx context.Context, info sys_model.PersonLicense, data *T) (response *T, err error) {
	newData := &sys_entity.SysPersonLicense{}
	gconv.Struct(data, newData)

	{
		userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

		// 用户资源文件夹
		userFolder := "./resources/license/" + gconv.String(newData.Id)
		fileAt := gtime.Now().Format("YmdHis")
		if !gfile.Exists(info.IdcardFrontPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardFrontPath), userId, "请上传身份证头像面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证头像面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/front_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.IdcardFrontPath = fileInfo.Src
		}

		if !gfile.Exists(info.IdcardBackPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardBackPath), userId, "请上传身份证国徽面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证国徽面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/back_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.IdcardBackPath = fileInfo.Src
		}

	}

	gconv.Struct(newData, data)
	return data, err
}
