// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	ISdkBaidu interface {
		GetBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.BaiduSdkConfToken, err error)
		GetBaiduSdkConfList(ctx context.Context) ([]*sys_model.BaiduSdkConf, error)
		GetBaiduSdkConf(ctx context.Context, identifier string) (*sys_model.BaiduSdkConf, error)
		SaveBaiduSdkConf(ctx context.Context, info *sys_model.BaiduSdkConf, isCreate bool) (*sys_model.BaiduSdkConf, error)
		DeleteBaiduSdkConf(ctx context.Context, identifier string) (bool, error)
		OCRBankCard(ctx context.Context, imageBase64 string) (*sys_model.OCRBankCard, error)
		OCRIDCard(ctx context.Context, imageBase64 string, detectRisk string, idCardSide string) (*sys_model.BaiduSdkOCRIDCard, error)
		OCRBusinessLicense(ctx context.Context, imageBase64 string) (*sys_model.BusinessLicenseOCR, error)
	}
)

var (
	localSdkBaidu ISdkBaidu
)

func SdkBaidu() ISdkBaidu {
	if localSdkBaidu == nil {
		panic("implement not found for interface ISdkBaidu, forgot register?")
	}
	return localSdkBaidu
}

func RegisterSdkBaidu(i ISdkBaidu) {
	localSdkBaidu = i
}
