package model

import (
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyUpload "github.com/SupenBysz/gf-admin-community/model/enum/upload"
	"github.com/gogf/gf/v2/net/ghttp"
)

// FileUploadInput 上传文件输入参数
type FileUploadInput struct {
	File       *ghttp.UploadFile `json:"file" type:"file" dc:"请选择文件，以form-data方式提交"` // 上传文件对象
	Name       string            `json:"name" dc:"自定义文件名"`                           // 自定义文件名称
	RandomName bool              `json:"randomName" dc:"是否随机文件名" default:"true"`     // 是否随机命名文件
}

// FileUploadOutput 上传文件返回参数
type FileUploadOutput struct {
	Id     int64  `json:"id" dc:"ID"`           // 数据表ID
	Name   string `json:"name" dc:"文件名称"`       // 文件名称
	Path   string `json:"path" dc:"文件路径"`       // 本地路径
	Url    string `json:"url" dc:"文件访问地址"`      // 访问URL，可能只是URI
	UserId int64  `json:"userId" dc:"文件所属用户ID"` // 用户ID
	Size   int64  `json:"size" dc:"文件大小"`       // 文件大小
}

// OCRIDCardFileUploadInput 上传身份证请求参数
type OCRIDCardFileUploadInput struct {
	FileUploadInput
	DetectRisk string `json:"detectRisk" default:"false" v:"in:true,false" dc:"是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)检测功能，默认不开启"`
	IDCardSide string `json:"idCardSide" default:"front" v:"in:front,back#参数错误|正反面参数错误" dc:"front身份证含照片的一面,back身份证带国徽的一面,自动检测身份证正反面，如果传参指定方向与图片相反，支持正常识别，返回参数image_status字段为 reversed_side "`
}

// IDCardWithOCR 身份证识别响应信息
type IDCardWithOCR struct {
	FileUploadOutput
	BaiduSdkOCRIDCard
	// OCRBusinessLicense OCRBusinessLicense `json:"orcBusinessLicense" dc:"营业执照识别的信息"`
}

// OCRBusinessLicense 上传新版营业执照
type OCRBusinessLicense struct {
	FileUploadInput
}

// BusinessLicenseWithOCR 营业执照识别响应信息
type BusinessLicenseWithOCR struct {
	FileUploadOutput
	BusinessLicenseOCR
}

// BankCardWithOCRInput 上传银行卡请求参数
type BankCardWithOCRInput struct {
	FileUploadInput
}

// BankCardWithOCR 银行卡识别响应信息
type BankCardWithOCR struct {
	// 上传文件返回数据
	FileUploadOutput

	// SDK返回的识别数据
	BaiduSdkOCRBankCard
}

type FileHookFunc HookFunc[kyUpload.EventStateEnum, entity.SysFile]
type FileHookInfo HookEventType[kyUpload.EventStateEnum, FileHookFunc]
