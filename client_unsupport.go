//go:build !linux || !cgo
// +build !linux !cgo

package WeWorkFinanceSDK

import (
	"fmt"
)

// NewClient 初始化函数
// Return值=0表示该API调用成功
//
// @param [in]  sdk			NewSdk返回的sdk指针
// @param [in]  corpId       调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
// @param [in]  secret		 聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
// @param [in]  privateKey    消息加密私钥，可以在企业微信管理端--管理工具--消息加密公钥查看对用公钥，私钥一般由自己保存
//
// @return 返回是否初始化成功
//      0   - 成功
//  	!=0 - 失败
//
func NewClient(corpId string, corpSecret string, rsaPrivateKey string) (Client, error) {
	return nil, fmt.Errorf("该 SDK 目前只支持 Linux 平台运行，请在 Linux 环境下安装并执行")
}
