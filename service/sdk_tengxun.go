// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
)

type (
	ISdkTengxun interface {
		GetTengxunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.TengxunSdkConfToken, err error)
		GetTengxunSdkConfList(ctx context.Context) (*[]model.TengxunSdkConf, error)
		GetTengxunSdkConf(ctx context.Context, identifier string) (tokenInfo *model.TengxunSdkConf, err error)
		SaveTengxunSdkConf(ctx context.Context, info model.TengxunSdkConf, isCreate bool) (*model.TengxunSdkConf, error)
		DeleteTengxunSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkTengxun ISdkTengxun
)

func SdkTengxun() ISdkTengxun {
	if localSdkTengxun == nil {
		panic("implement not found for interface ISdkTengxun, forgot register?")
	}
	return localSdkTengxun
}

func RegisterSdkTengxun(i ISdkTengxun) {
	localSdkTengxun = i
}
