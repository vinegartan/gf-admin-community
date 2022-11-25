// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/SupenBysz/gf-admin-community/model"
)

type (
	IBizCtx interface {
		Init(r *ghttp.Request, customCtx *model.Context)
		Get(ctx context.Context) *model.Context
		SetUser(ctx context.Context, claimsUser *model.JwtCustomClaims)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
