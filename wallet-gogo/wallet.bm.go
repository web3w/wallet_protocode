// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: wallet.proto

package pbwallet

import (
	"context"

	bm "github.com/gisvr/golib/net/http/blademaster"
	"github.com/gisvr/golib/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathWalletAddWallet = "/wallet.api.Wallet/AddWallet"
var PathWalletGetWallet = "/wallet.api.Wallet/GetWallet"
var PathWalletUpdateWallet = "/wallet.api.Wallet/UpdateWallet"
var PathWalletRemoveWallet = "/wallet.api.Wallet/RemoveWallet"
var PathWalletListWallet = "/wallet.api.Wallet/ListWallet"
var PathWalletListCoinInfo = "/wallet.api.Wallet/ListCoinInfo"
var PathWalletAddWalletCoin = "/wallet.api.Wallet/AddWalletCoin"
var PathWalletRemoveWalletCoin = "/wallet.api.Wallet/RemoveWalletCoin"
var PathWalletListWalletCoin = "/wallet.api.Wallet/ListWalletCoin"
var PathWalletGetWalletCoin = "/wallet.api.Wallet/GetWalletCoin"
var PathWalletListWalletDetail = "/wallet.api.Wallet/ListWalletDetail"
var PathWalletListWalletCoinDetail = "/wallet.api.Wallet/ListWalletCoinDetail"
var PathWalletGetNewAddress = "/wallet.api.Wallet/GetNewAddress"
var PathWalletHideAddress = "/wallet.api.Wallet/HideAddress"
var PathWalletListAddress = "/wallet.api.Wallet/ListAddress"
var PathWalletHasAddress = "/wallet.api.Wallet/HasAddress"
var PathWalletCheckAddress = "/wallet.api.Wallet/CheckAddress"
var PathWalletGetCoinFee = "/wallet.api.Wallet/GetCoinFee"
var PathWalletListCoinTx = "/wallet.api.Wallet/ListCoinTx"
var PathWalletQueryCoinTx = "/wallet.api.Wallet/QueryCoinTx"
var PathWalletNewWithdraw = "/wallet.api.Wallet/NewWithdraw"
var PathWalletGetWithdrawDetail = "/wallet.api.Wallet/GetWithdrawDetail"
var PathWalletCollectCallback = "/wallet.api.Wallet/CollectCallback"
var PathWalletImportOldAddress = "/wallet.api.Wallet/ImportOldAddress"
var PathWalletGetWalletAsset = "/wallet.api.Wallet/GetWalletAsset"
var PathWalletListWalletAsset = "/wallet.api.Wallet/ListWalletAsset"
var PathWalletAddWithdrawSetting = "/wallet.api.Wallet/AddWithdrawSetting"
var PathWalletUpdateWithdrawSetting = "/wallet.api.Wallet/UpdateWithdrawSetting"
var PathWalletRemoveWithdrawSetting = "/wallet.api.Wallet/RemoveWithdrawSetting"
var PathWalletGetWithdrawSetting = "/wallet.api.Wallet/GetWithdrawSetting"
var PathWalletAddWithdrawQuota = "/wallet.api.Wallet/AddWithdrawQuota"
var PathWalletRemoveWithdrawQuota = "/wallet.api.Wallet/RemoveWithdrawQuota"
var PathWalletListWithdrawQuota = "/wallet.api.Wallet/ListWithdrawQuota"
var PathWalletAddWithdrawWhitelist = "/wallet.api.Wallet/AddWithdrawWhitelist"
var PathWalletRemoveWithdrawWhitelist = "/wallet.api.Wallet/RemoveWithdrawWhitelist"
var PathWalletListWithdrawWhitelist = "/wallet.api.Wallet/ListWithdrawWhitelist"
var PathWalletUpdateWithdrawPolicy = "/wallet.api.Wallet/UpdateWithdrawPolicy"
var PathWalletListWithdrawPolicy = "/wallet.api.Wallet/ListWithdrawPolicy"
var PathWalletGetWithdrawPolicy = "/wallet.api.Wallet/GetWithdrawPolicy"
var PathWalletAddMessage = "/wallet.api.Wallet/AddMessage"
var PathWalletUpdateMessage = "/wallet.api.Wallet/UpdateMessage"
var PathWalletListMessage = "/wallet.api.Wallet/ListMessage"

const (
	PermissionAddWallet               = "PermissionWalletAddWallet"
	PermissionGetWallet               = "PermissionWalletGetWallet"
	PermissionUpdateWallet            = "PermissionWalletUpdateWallet"
	PermissionRemoveWallet            = "PermissionWalletRemoveWallet"
	PermissionListWallet              = "PermissionWalletListWallet"
	PermissionListCoinInfo            = "PermissionWalletListCoinInfo"
	PermissionAddWalletCoin           = "PermissionWalletAddWalletCoin"
	PermissionRemoveWalletCoin        = "PermissionWalletRemoveWalletCoin"
	PermissionListWalletCoin          = "PermissionWalletListWalletCoin"
	PermissionGetWalletCoin           = "PermissionWalletGetWalletCoin"
	PermissionListWalletDetail        = "PermissionWalletListWalletDetail"
	PermissionListWalletCoinDetail    = "PermissionWalletListWalletCoinDetail"
	PermissionGetNewAddress           = "PermissionWalletGetNewAddress"
	PermissionHideAddress             = "PermissionWalletHideAddress"
	PermissionListAddress             = "PermissionWalletListAddress"
	PermissionHasAddress              = "PermissionWalletHasAddress"
	PermissionCheckAddress            = "PermissionWalletCheckAddress"
	PermissionGetCoinFee              = "PermissionWalletGetCoinFee"
	PermissionListCoinTx              = "PermissionWalletListCoinTx"
	PermissionQueryCoinTx             = "PermissionWalletQueryCoinTx"
	PermissionNewWithdraw             = "PermissionWalletNewWithdraw"
	PermissionGetWithdrawDetail       = "PermissionWalletGetWithdrawDetail"
	PermissionCollectCallback         = "PermissionWalletCollectCallback"
	PermissionImportOldAddress        = "PermissionWalletImportOldAddress"
	PermissionGetWalletAsset          = "PermissionWalletGetWalletAsset"
	PermissionListWalletAsset         = "PermissionWalletListWalletAsset"
	PermissionAddWithdrawSetting      = "PermissionWalletAddWithdrawSetting"
	PermissionUpdateWithdrawSetting   = "PermissionWalletUpdateWithdrawSetting"
	PermissionRemoveWithdrawSetting   = "PermissionWalletRemoveWithdrawSetting"
	PermissionGetWithdrawSetting      = "PermissionWalletGetWithdrawSetting"
	PermissionAddWithdrawQuota        = "PermissionWalletAddWithdrawQuota"
	PermissionRemoveWithdrawQuota     = "PermissionWalletRemoveWithdrawQuota"
	PermissionListWithdrawQuota       = "PermissionWalletListWithdrawQuota"
	PermissionAddWithdrawWhitelist    = "PermissionWalletAddWithdrawWhitelist"
	PermissionRemoveWithdrawWhitelist = "PermissionWalletRemoveWithdrawWhitelist"
	PermissionListWithdrawWhitelist   = "PermissionWalletListWithdrawWhitelist"
	PermissionUpdateWithdrawPolicy    = "PermissionWalletUpdateWithdrawPolicy"
	PermissionListWithdrawPolicy      = "PermissionWalletListWithdrawPolicy"
	PermissionGetWithdrawPolicy       = "PermissionWalletGetWithdrawPolicy"
	PermissionAddMessage              = "PermissionWalletAddMessage"
	PermissionUpdateMessage           = "PermissionWalletUpdateMessage"
	PermissionListMessage             = "PermissionWalletListMessage"
)

type Permission struct {
	Module      string
	Name        string
	Url         string
	Description string
}

var Perms = []Permission{
	Permission{"Wallet", PermissionAddWallet, PathWalletAddWallet, "新增钱包"},
	Permission{"Wallet", PermissionGetWallet, PathWalletGetWallet, " 根据钱包ID获取钱包信息。"},
	Permission{"Wallet", PermissionUpdateWallet, PathWalletUpdateWallet, "更新钱包"},
	Permission{"Wallet", PermissionRemoveWallet, PathWalletRemoveWallet, "移除钱包"},
	Permission{"Wallet", PermissionListWallet, PathWalletListWallet, "获取钱包信息"},
	Permission{"Wallet", PermissionListCoinInfo, PathWalletListCoinInfo, "获取支持的钱包"},
	Permission{"Wallet", PermissionAddWalletCoin, PathWalletAddWalletCoin, "添加coin钱包"},
	Permission{"Wallet", PermissionRemoveWalletCoin, PathWalletRemoveWalletCoin, "删除Coin钱包"},
	Permission{"Wallet", PermissionListWalletCoin, PathWalletListWalletCoin, ""},
	Permission{"Wallet", PermissionGetWalletCoin, PathWalletGetWalletCoin, " 根据ID获取WalletCoin信息。"},
	Permission{"Wallet", PermissionListWalletDetail, PathWalletListWalletDetail, "获取钱包的详情，包括添加的币种、资产情况等"},
	Permission{"Wallet", PermissionListWalletCoinDetail, PathWalletListWalletCoinDetail, "获取钱包coin及资产列表"},
	Permission{"Wallet", PermissionGetNewAddress, PathWalletGetNewAddress, "增加新地址并加入扫块监控"},
	Permission{"Wallet", PermissionHideAddress, PathWalletHideAddress, "隐藏地址"},
	Permission{"Wallet", PermissionListAddress, PathWalletListAddress, "获取指定钱包coin的所有地址"},
	Permission{"Wallet", PermissionHasAddress, PathWalletHasAddress, "验证地址归属 即地址存在否"},
	Permission{"Wallet", PermissionCheckAddress, PathWalletCheckAddress, "验证地址合法性"},
	Permission{"Wallet", PermissionGetCoinFee, PathWalletGetCoinFee, "获取coin当前的最佳手续费率"},
	Permission{"Wallet", PermissionListCoinTx, PathWalletListCoinTx, "获取CoinTx列表"},
	Permission{"Wallet", PermissionQueryCoinTx, PathWalletQueryCoinTx, "查询CoinTx"},
	Permission{"Wallet", PermissionNewWithdraw, PathWalletNewWithdraw, "申请提现"},
	Permission{"Wallet", PermissionGetWithdrawDetail, PathWalletGetWithdrawDetail, "获取交易详细信息"},
	Permission{"Wallet", PermissionCollectCallback, PathWalletCollectCallback, "外部做归集时，提现完成回调通知"},
	Permission{"Wallet", PermissionImportOldAddress, PathWalletImportOldAddress, "导入老地址"},
	Permission{"Wallet", PermissionGetWalletAsset, PathWalletGetWalletAsset, "获取指定coin钱包信息及余额"},
	Permission{"Wallet", PermissionListWalletAsset, PathWalletListWalletAsset, "获取账户指定钱包所有coin的信息及余额"},
	Permission{"Wallet", PermissionAddWithdrawSetting, PathWalletAddWithdrawSetting, "提现策略设置"},
	Permission{"Wallet", PermissionUpdateWithdrawSetting, PathWalletUpdateWithdrawSetting, ""},
	Permission{"Wallet", PermissionRemoveWithdrawSetting, PathWalletRemoveWithdrawSetting, ""},
	Permission{"Wallet", PermissionGetWithdrawSetting, PathWalletGetWithdrawSetting, ""},
	Permission{"Wallet", PermissionAddWithdrawQuota, PathWalletAddWithdrawQuota, "提现限额"},
	Permission{"Wallet", PermissionRemoveWithdrawQuota, PathWalletRemoveWithdrawQuota, ""},
	Permission{"Wallet", PermissionListWithdrawQuota, PathWalletListWithdrawQuota, ""},
	Permission{"Wallet", PermissionAddWithdrawWhitelist, PathWalletAddWithdrawWhitelist, "提现白名单"},
	Permission{"Wallet", PermissionRemoveWithdrawWhitelist, PathWalletRemoveWithdrawWhitelist, ""},
	Permission{"Wallet", PermissionListWithdrawWhitelist, PathWalletListWithdrawWhitelist, ""},
	Permission{"Wallet", PermissionUpdateWithdrawPolicy, PathWalletUpdateWithdrawPolicy, ""},
	Permission{"Wallet", PermissionListWithdrawPolicy, PathWalletListWithdrawPolicy, ""},
	Permission{"Wallet", PermissionGetWithdrawPolicy, PathWalletGetWithdrawPolicy, ""},
	Permission{"Wallet", PermissionAddMessage, PathWalletAddMessage, "推送历史记录"},
	Permission{"Wallet", PermissionUpdateMessage, PathWalletUpdateMessage, ""},
	Permission{"Wallet", PermissionListMessage, PathWalletListMessage, ""},
}

// WalletBMServer is the server API for Wallet service.
type WalletBMServer interface {
	// 新增钱包
	AddWallet(ctx context.Context, req *AddWalletRequest) (resp *AddWalletResponse, err error)

	// 根据钱包ID获取钱包信息。
	GetWallet(ctx context.Context, req *GetWalletRequest) (resp *GetWalletResponse, err error)

	// 更新钱包
	UpdateWallet(ctx context.Context, req *UpdateWalletRequest) (resp *UpdateWalletResponse, err error)

	// 移除钱包
	RemoveWallet(ctx context.Context, req *RemoveWalletRequest) (resp *RemoveWalletResponse, err error)

	// 获取钱包信息
	ListWallet(ctx context.Context, req *ListWalletRequest) (resp *ListWalletResponse, err error)

	// 获取支持的钱包
	ListCoinInfo(ctx context.Context, req *ListCoinInfoRequest) (resp *ListCoinInfoResponse, err error)

	// 添加coin钱包
	AddWalletCoin(ctx context.Context, req *AddWalletCoinRequest) (resp *AddWalletCoinResponse, err error)

	// 删除Coin钱包
	RemoveWalletCoin(ctx context.Context, req *RemoveWalletCoinRequest) (resp *RemoveWalletCoinResponse, err error)

	ListWalletCoin(ctx context.Context, req *ListWalletCoinRequest) (resp *ListWalletCoinResponse, err error)

	// 根据ID获取WalletCoin信息。
	GetWalletCoin(ctx context.Context, req *GetWalletCoinRequest) (resp *GetWalletCoinResponse, err error)

	// 获取钱包的详情，包括添加的币种、资产情况等
	ListWalletDetail(ctx context.Context, req *ListWalletDetailRequest) (resp *ListWalletDetailResponse, err error)

	// 获取钱包coin及资产列表
	ListWalletCoinDetail(ctx context.Context, req *ListWalletCoinDetailRequest) (resp *ListWalletCoinDetailResponse, err error)

	// 增加新地址并加入扫块监控
	GetNewAddress(ctx context.Context, req *GetNewAddressRequest) (resp *GetNewAddressResponse, err error)

	// 隐藏地址
	HideAddress(ctx context.Context, req *HideAddressRequest) (resp *HideAddressResponse, err error)

	// 获取指定钱包coin的所有地址
	ListAddress(ctx context.Context, req *ListAddressRequest) (resp *ListAddressResponse, err error)

	// 验证地址归属 即地址存在否
	HasAddress(ctx context.Context, req *HasAddressRequest) (resp *HasAddressResponse, err error)

	// 验证地址合法性
	CheckAddress(ctx context.Context, req *CheckAddressRequest) (resp *CheckAddressResponse, err error)

	// 获取coin当前的最佳手续费率
	GetCoinFee(ctx context.Context, req *GetCoinFeeRequest) (resp *GetCoinFeeResponse, err error)

	// 获取CoinTx列表
	ListCoinTx(ctx context.Context, req *ListCoinTxRequest) (resp *ListCoinTxResponse, err error)

	// 查询CoinTx
	QueryCoinTx(ctx context.Context, req *QueryCoinTxRequest) (resp *QueryCoinTxResponse, err error)

	// 申请提现
	NewWithdraw(ctx context.Context, req *NewWithdrawRequest) (resp *NewWithdrawResponse, err error)

	// 获取交易详细信息
	GetWithdrawDetail(ctx context.Context, req *GetWithdrawDetailRequest) (resp *GetWithdrawDetailResponse, err error)

	// 外部做归集时，提现完成回调通知
	CollectCallback(ctx context.Context, req *CollectCallbackRequest) (resp *CollectCallbackResponse, err error)

	// 导入老地址
	ImportOldAddress(ctx context.Context, req *ImportOldAddressRequest) (resp *ImportOldAddressResponse, err error)

	// 获取指定coin钱包信息及余额
	GetWalletAsset(ctx context.Context, req *GetWalletAssetRequest) (resp *GetWalletAssetResponse, err error)

	// 获取账户指定钱包所有coin的信息及余额
	ListWalletAsset(ctx context.Context, req *ListWalletAssetRequest) (resp *ListWalletAssetResponse, err error)

	// 提现策略设置
	AddWithdrawSetting(ctx context.Context, req *AddWithdrawSettingRequest) (resp *AddWithdrawSettingResponse, err error)

	UpdateWithdrawSetting(ctx context.Context, req *UpdateWithdrawSettingRequest) (resp *UpdateWithdrawSettingResponse, err error)

	RemoveWithdrawSetting(ctx context.Context, req *RemoveWithdrawSettingRequest) (resp *RemoveWithdrawSettingResponse, err error)

	GetWithdrawSetting(ctx context.Context, req *GetWithdrawSettingRequest) (resp *GetWithdrawSettingResponse, err error)

	// 提现限额
	AddWithdrawQuota(ctx context.Context, req *AddWithdrawQuotaRequest) (resp *AddWithdrawQuotaResponse, err error)

	RemoveWithdrawQuota(ctx context.Context, req *RemoveWithdrawQuotaRequest) (resp *RemoveWithdrawQuotaResponse, err error)

	ListWithdrawQuota(ctx context.Context, req *ListWithdrawQuotaRequest) (resp *ListWithdrawQuotaResponse, err error)

	// 提现白名单
	AddWithdrawWhitelist(ctx context.Context, req *AddWithdrawWhitelistRequest) (resp *AddWithdrawWhitelistResponse, err error)

	RemoveWithdrawWhitelist(ctx context.Context, req *RemoveWithdrawWhitelistRequest) (resp *RemoveWithdrawWhitelistResponse, err error)

	ListWithdrawWhitelist(ctx context.Context, req *ListWithdrawWhitelistRequest) (resp *ListWithdrawWhitelistResponse, err error)

	UpdateWithdrawPolicy(ctx context.Context, req *UpdateWithdrawPolicyRequest) (resp *UpdateWithdrawPolicyResponse, err error)

	ListWithdrawPolicy(ctx context.Context, req *ListWithdrawPolicyRequest) (resp *ListWithdrawPolicyResponse, err error)

	GetWithdrawPolicy(ctx context.Context, req *GetWithdrawPolicyRequest) (resp *GetWithdrawPolicyResponse, err error)

	// 推送历史记录
	AddMessage(ctx context.Context, req *AddMessageRequest) (resp *AddMessageResponse, err error)

	UpdateMessage(ctx context.Context, req *UpdateMessageRequest) (resp *UpdateMessageResponse, err error)

	ListMessage(ctx context.Context, req *ListMessageRequest) (resp *ListMessageResponse, err error)
}

var WalletSvc WalletBMServer

type JSONReaderFunc func(c *bm.Context, data interface{}, err error)

func DefaultJSONReader(c *bm.Context, data interface{}, err error) {
	c.JSON(data, err)
}

var JF JSONReaderFunc = DefaultJSONReader

func WalletAddWallet(c *bm.Context) {
	p := new(AddWalletRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddWallet(c, p)
	JF(c, resp, err)
}

func WalletGetWallet(c *bm.Context) {
	p := new(GetWalletRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWallet(c, p)
	JF(c, resp, err)
}

func WalletUpdateWallet(c *bm.Context) {
	p := new(UpdateWalletRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.UpdateWallet(c, p)
	JF(c, resp, err)
}

func WalletRemoveWallet(c *bm.Context) {
	p := new(RemoveWalletRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.RemoveWallet(c, p)
	JF(c, resp, err)
}

func WalletListWallet(c *bm.Context) {
	p := new(ListWalletRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWallet(c, p)
	JF(c, resp, err)
}

func WalletListCoinInfo(c *bm.Context) {
	p := new(ListCoinInfoRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListCoinInfo(c, p)
	JF(c, resp, err)
}

func WalletAddWalletCoin(c *bm.Context) {
	p := new(AddWalletCoinRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddWalletCoin(c, p)
	JF(c, resp, err)
}

func WalletRemoveWalletCoin(c *bm.Context) {
	p := new(RemoveWalletCoinRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.RemoveWalletCoin(c, p)
	JF(c, resp, err)
}

func WalletListWalletCoin(c *bm.Context) {
	p := new(ListWalletCoinRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWalletCoin(c, p)
	JF(c, resp, err)
}

func WalletGetWalletCoin(c *bm.Context) {
	p := new(GetWalletCoinRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWalletCoin(c, p)
	JF(c, resp, err)
}

func WalletListWalletDetail(c *bm.Context) {
	p := new(ListWalletDetailRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWalletDetail(c, p)
	JF(c, resp, err)
}

func WalletListWalletCoinDetail(c *bm.Context) {
	p := new(ListWalletCoinDetailRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWalletCoinDetail(c, p)
	JF(c, resp, err)
}

func WalletGetNewAddress(c *bm.Context) {
	p := new(GetNewAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetNewAddress(c, p)
	JF(c, resp, err)
}

func WalletHideAddress(c *bm.Context) {
	p := new(HideAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.HideAddress(c, p)
	JF(c, resp, err)
}

func WalletListAddress(c *bm.Context) {
	p := new(ListAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListAddress(c, p)
	JF(c, resp, err)
}

func WalletHasAddress(c *bm.Context) {
	p := new(HasAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.HasAddress(c, p)
	JF(c, resp, err)
}

func WalletCheckAddress(c *bm.Context) {
	p := new(CheckAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.CheckAddress(c, p)
	JF(c, resp, err)
}

func WalletGetCoinFee(c *bm.Context) {
	p := new(GetCoinFeeRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetCoinFee(c, p)
	JF(c, resp, err)
}

func WalletListCoinTx(c *bm.Context) {
	p := new(ListCoinTxRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListCoinTx(c, p)
	JF(c, resp, err)
}

func WalletQueryCoinTx(c *bm.Context) {
	p := new(QueryCoinTxRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.QueryCoinTx(c, p)
	JF(c, resp, err)
}

func WalletNewWithdraw(c *bm.Context) {
	p := new(NewWithdrawRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.NewWithdraw(c, p)
	JF(c, resp, err)
}

func WalletGetWithdrawDetail(c *bm.Context) {
	p := new(GetWithdrawDetailRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWithdrawDetail(c, p)
	JF(c, resp, err)
}

func WalletCollectCallback(c *bm.Context) {
	p := new(CollectCallbackRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.CollectCallback(c, p)
	JF(c, resp, err)
}

func WalletImportOldAddress(c *bm.Context) {
	p := new(ImportOldAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ImportOldAddress(c, p)
	JF(c, resp, err)
}

func WalletGetWalletAsset(c *bm.Context) {
	p := new(GetWalletAssetRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWalletAsset(c, p)
	JF(c, resp, err)
}

func WalletListWalletAsset(c *bm.Context) {
	p := new(ListWalletAssetRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWalletAsset(c, p)
	JF(c, resp, err)
}

func WalletAddWithdrawSetting(c *bm.Context) {
	p := new(AddWithdrawSettingRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddWithdrawSetting(c, p)
	JF(c, resp, err)
}

func WalletUpdateWithdrawSetting(c *bm.Context) {
	p := new(UpdateWithdrawSettingRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.UpdateWithdrawSetting(c, p)
	JF(c, resp, err)
}

func WalletRemoveWithdrawSetting(c *bm.Context) {
	p := new(RemoveWithdrawSettingRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.RemoveWithdrawSetting(c, p)
	JF(c, resp, err)
}

func WalletGetWithdrawSetting(c *bm.Context) {
	p := new(GetWithdrawSettingRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWithdrawSetting(c, p)
	JF(c, resp, err)
}

func WalletAddWithdrawQuota(c *bm.Context) {
	p := new(AddWithdrawQuotaRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddWithdrawQuota(c, p)
	JF(c, resp, err)
}

func WalletRemoveWithdrawQuota(c *bm.Context) {
	p := new(RemoveWithdrawQuotaRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.RemoveWithdrawQuota(c, p)
	JF(c, resp, err)
}

func WalletListWithdrawQuota(c *bm.Context) {
	p := new(ListWithdrawQuotaRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWithdrawQuota(c, p)
	JF(c, resp, err)
}

func WalletAddWithdrawWhitelist(c *bm.Context) {
	p := new(AddWithdrawWhitelistRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddWithdrawWhitelist(c, p)
	JF(c, resp, err)
}

func WalletRemoveWithdrawWhitelist(c *bm.Context) {
	p := new(RemoveWithdrawWhitelistRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.RemoveWithdrawWhitelist(c, p)
	JF(c, resp, err)
}

func WalletListWithdrawWhitelist(c *bm.Context) {
	p := new(ListWithdrawWhitelistRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWithdrawWhitelist(c, p)
	JF(c, resp, err)
}

func WalletUpdateWithdrawPolicy(c *bm.Context) {
	p := new(UpdateWithdrawPolicyRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.UpdateWithdrawPolicy(c, p)
	JF(c, resp, err)
}

func WalletListWithdrawPolicy(c *bm.Context) {
	p := new(ListWithdrawPolicyRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListWithdrawPolicy(c, p)
	JF(c, resp, err)
}

func WalletGetWithdrawPolicy(c *bm.Context) {
	p := new(GetWithdrawPolicyRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.GetWithdrawPolicy(c, p)
	JF(c, resp, err)
}

func WalletAddMessage(c *bm.Context) {
	p := new(AddMessageRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.AddMessage(c, p)
	JF(c, resp, err)
}

func WalletUpdateMessage(c *bm.Context) {
	p := new(UpdateMessageRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.UpdateMessage(c, p)
	JF(c, resp, err)
}

func WalletListMessage(c *bm.Context) {
	p := new(ListMessageRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WalletSvc.ListMessage(c, p)
	JF(c, resp, err)
}

// RegisterWalletBMServer Register the blademaster route
func RegisterWalletBMServer(e *bm.Engine, server WalletBMServer) {
	WalletSvc = server
	e.GET(PathWalletAddWallet, WalletAddWallet)
	e.GET(PathWalletGetWallet, WalletGetWallet)
	e.GET(PathWalletUpdateWallet, WalletUpdateWallet)
	e.GET(PathWalletRemoveWallet, WalletRemoveWallet)
	e.GET(PathWalletListWallet, WalletListWallet)
	e.GET(PathWalletListCoinInfo, WalletListCoinInfo)
	e.GET(PathWalletAddWalletCoin, WalletAddWalletCoin)
	e.GET(PathWalletRemoveWalletCoin, WalletRemoveWalletCoin)
	e.GET(PathWalletListWalletCoin, WalletListWalletCoin)
	e.GET(PathWalletGetWalletCoin, WalletGetWalletCoin)
	e.GET(PathWalletListWalletDetail, WalletListWalletDetail)
	e.GET(PathWalletListWalletCoinDetail, WalletListWalletCoinDetail)
	e.GET(PathWalletGetNewAddress, WalletGetNewAddress)
	e.GET(PathWalletHideAddress, WalletHideAddress)
	e.GET(PathWalletListAddress, WalletListAddress)
	e.GET(PathWalletHasAddress, WalletHasAddress)
	e.GET(PathWalletCheckAddress, WalletCheckAddress)
	e.GET(PathWalletGetCoinFee, WalletGetCoinFee)
	e.GET(PathWalletListCoinTx, WalletListCoinTx)
	e.GET(PathWalletQueryCoinTx, WalletQueryCoinTx)
	e.GET(PathWalletNewWithdraw, WalletNewWithdraw)
	e.GET(PathWalletGetWithdrawDetail, WalletGetWithdrawDetail)
	e.GET(PathWalletCollectCallback, WalletCollectCallback)
	e.GET(PathWalletImportOldAddress, WalletImportOldAddress)
	e.GET(PathWalletGetWalletAsset, WalletGetWalletAsset)
	e.GET(PathWalletListWalletAsset, WalletListWalletAsset)
	e.GET(PathWalletAddWithdrawSetting, WalletAddWithdrawSetting)
	e.GET(PathWalletUpdateWithdrawSetting, WalletUpdateWithdrawSetting)
	e.GET(PathWalletRemoveWithdrawSetting, WalletRemoveWithdrawSetting)
	e.GET(PathWalletGetWithdrawSetting, WalletGetWithdrawSetting)
	e.GET(PathWalletAddWithdrawQuota, WalletAddWithdrawQuota)
	e.GET(PathWalletRemoveWithdrawQuota, WalletRemoveWithdrawQuota)
	e.GET(PathWalletListWithdrawQuota, WalletListWithdrawQuota)
	e.GET(PathWalletAddWithdrawWhitelist, WalletAddWithdrawWhitelist)
	e.GET(PathWalletRemoveWithdrawWhitelist, WalletRemoveWithdrawWhitelist)
	e.GET(PathWalletListWithdrawWhitelist, WalletListWithdrawWhitelist)
	e.GET(PathWalletUpdateWithdrawPolicy, WalletUpdateWithdrawPolicy)
	e.GET(PathWalletListWithdrawPolicy, WalletListWithdrawPolicy)
	e.GET(PathWalletGetWithdrawPolicy, WalletGetWithdrawPolicy)
	e.GET(PathWalletAddMessage, WalletAddMessage)
	e.GET(PathWalletUpdateMessage, WalletUpdateMessage)
	e.GET(PathWalletListMessage, WalletListMessage)
}
