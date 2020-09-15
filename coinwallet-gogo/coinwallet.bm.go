// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: coinwallet.proto

package pbcoinwallet

import (
	"context"

	bm "github.com/gisvr/golib/net/http/blademaster"
	"github.com/gisvr/golib/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathCoinWalletGetBlockInfo = "/coinwallet.api.CoinWallet/GetBlockInfo"
var PathCoinWalletGetTxinfo = "/coinwallet.api.CoinWallet/GetTxinfo"
var PathCoinWalletGetNewAddress = "/coinwallet.api.CoinWallet/GetNewAddress"
var PathCoinWalletValidateAddress = "/coinwallet.api.CoinWallet/ValidateAddress"
var PathCoinWalletSendTo = "/coinwallet.api.CoinWallet/SendTo"
var PathCoinWalletSendToMany = "/coinwallet.api.CoinWallet/SendToMany"
var PathCoinWalletGetBalanceByAddress = "/coinwallet.api.CoinWallet/GetBalanceByAddress"
var PathCoinWalletGetBalance = "/coinwallet.api.CoinWallet/GetBalance"
var PathCoinWalletAddMonitorAddress = "/coinwallet.api.CoinWallet/AddMonitorAddress"
var PathCoinWalletRemoveMonitorAddress = "/coinwallet.api.CoinWallet/RemoveMonitorAddress"
var PathCoinWalletSetCallback = "/coinwallet.api.CoinWallet/SetCallback"
var PathCoinWalletGetEstimateFee = "/coinwallet.api.CoinWallet/GetEstimateFee"
var PathCoinWalletGetMnemonicFingerprinting = "/coinwallet.api.CoinWallet/GetMnemonicFingerprinting"
var PathCoinWalletImportPrivateKey = "/coinwallet.api.CoinWallet/ImportPrivateKey"

const (
	PermissionGetBlockInfo              = "PermissionCoinWalletGetBlockInfo"
	PermissionGetTxinfo                 = "PermissionCoinWalletGetTxinfo"
	PermissionGetNewAddress             = "PermissionCoinWalletGetNewAddress"
	PermissionValidateAddress           = "PermissionCoinWalletValidateAddress"
	PermissionSendTo                    = "PermissionCoinWalletSendTo"
	PermissionSendToMany                = "PermissionCoinWalletSendToMany"
	PermissionGetBalanceByAddress       = "PermissionCoinWalletGetBalanceByAddress"
	PermissionGetBalance                = "PermissionCoinWalletGetBalance"
	PermissionAddMonitorAddress         = "PermissionCoinWalletAddMonitorAddress"
	PermissionRemoveMonitorAddress      = "PermissionCoinWalletRemoveMonitorAddress"
	PermissionSetCallback               = "PermissionCoinWalletSetCallback"
	PermissionGetEstimateFee            = "PermissionCoinWalletGetEstimateFee"
	PermissionGetMnemonicFingerprinting = "PermissionCoinWalletGetMnemonicFingerprinting"
	PermissionImportPrivateKey          = "PermissionCoinWalletImportPrivateKey"
)

type Permission struct {
	Module      string
	Name        string
	Url         string
	Description string
}

var Perms = []Permission{
	Permission{"CoinWallet", PermissionGetBlockInfo, PathCoinWalletGetBlockInfo, "获取最新区块信息"},
	Permission{"CoinWallet", PermissionGetTxinfo, PathCoinWalletGetTxinfo, "获取tx 信息"},
	Permission{"CoinWallet", PermissionGetNewAddress, PathCoinWalletGetNewAddress, "生成地址"},
	Permission{"CoinWallet", PermissionValidateAddress, PathCoinWalletValidateAddress, "验证地址"},
	Permission{"CoinWallet", PermissionSendTo, PathCoinWalletSendTo, "发币"},
	Permission{"CoinWallet", PermissionSendToMany, PathCoinWalletSendToMany, ""},
	Permission{"CoinWallet", PermissionGetBalanceByAddress, PathCoinWalletGetBalanceByAddress, "获取地址余额"},
	Permission{"CoinWallet", PermissionGetBalance, PathCoinWalletGetBalance, "获取钱包余额"},
	Permission{"CoinWallet", PermissionAddMonitorAddress, PathCoinWalletAddMonitorAddress, "添加钱包地址监控链上tx"},
	Permission{"CoinWallet", PermissionRemoveMonitorAddress, PathCoinWalletRemoveMonitorAddress, "移除钱包地址监控"},
	Permission{"CoinWallet", PermissionSetCallback, PathCoinWalletSetCallback, "监控事件回调"},
	Permission{"CoinWallet", PermissionGetEstimateFee, PathCoinWalletGetEstimateFee, "获取当前预估费率"},
	Permission{"CoinWallet", PermissionGetMnemonicFingerprinting, PathCoinWalletGetMnemonicFingerprinting, "钱包私钥特征，用于验证钱包内记录账目的地址和链上的地址是一致的"},
	Permission{"CoinWallet", PermissionImportPrivateKey, PathCoinWalletImportPrivateKey, "导入钱包并返回地址信息"},
}

// CoinWalletBMServer is the server API for CoinWallet service.
type CoinWalletBMServer interface {
	// 获取最新区块信息
	GetBlockInfo(ctx context.Context, req *GetBlockInfoRequest) (resp *GetBlockInfoResponse, err error)

	// 获取tx 信息
	GetTxinfo(ctx context.Context, req *GetTxinfoRequest) (resp *GetTxinfoResponse, err error)

	// 生成地址
	GetNewAddress(ctx context.Context, req *GetNewAddressRequest) (resp *GetNewAddressResponse, err error)

	// 验证地址
	ValidateAddress(ctx context.Context, req *ValidateAddressRequest) (resp *ValidateAddressResponse, err error)

	// 发币
	SendTo(ctx context.Context, req *SendToRequest) (resp *SendToResponse, err error)

	SendToMany(ctx context.Context, req *SendToManyRequest) (resp *SendToManyResponse, err error)

	// 获取地址余额
	GetBalanceByAddress(ctx context.Context, req *GetBalanceByAddressRequest) (resp *GetBalanceByAddressResponse, err error)

	// 获取钱包余额
	GetBalance(ctx context.Context, req *GetBalanceRequest) (resp *GetBalanceResponse, err error)

	// 添加钱包地址监控链上tx
	AddMonitorAddress(ctx context.Context, req *AddMonitorAddressRequest) (resp *AddMonitorAddressResponse, err error)

	// 移除钱包地址监控
	RemoveMonitorAddress(ctx context.Context, req *RemoveMonitorAddressRequest) (resp *RemoveMonitorAddressResponse, err error)

	// 监控事件回调
	SetCallback(ctx context.Context, req *SetCallbackRequest) (resp *SetCallbackResponse, err error)

	// 获取当前预估费率
	GetEstimateFee(ctx context.Context, req *GetEstimateFeeRequest) (resp *GetEstimateFeeResponse, err error)

	// 钱包私钥特征，用于验证钱包内记录账目的地址和链上的地址是一致的
	GetMnemonicFingerprinting(ctx context.Context, req *GetMnemonicFingerprintingRequest) (resp *GetMnemonicFingerprintingResponse, err error)

	// 导入钱包并返回地址信息
	ImportPrivateKey(ctx context.Context, req *ImportPrivateKeyRequest) (resp *ImportPrivateKeyResponse, err error)
}

var CoinWalletSvc CoinWalletBMServer

type JSONReaderFunc func(c *bm.Context, data interface{}, err error)

func DefaultJSONReader(c *bm.Context, data interface{}, err error) {
	c.JSON(data, err)
}

var JF JSONReaderFunc = DefaultJSONReader

func CoinWalletGetBlockInfo(c *bm.Context) {
	p := new(GetBlockInfoRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetBlockInfo(c, p)
	JF(c, resp, err)
}

func CoinWalletGetTxinfo(c *bm.Context) {
	p := new(GetTxinfoRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetTxinfo(c, p)
	JF(c, resp, err)
}

func CoinWalletGetNewAddress(c *bm.Context) {
	p := new(GetNewAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetNewAddress(c, p)
	JF(c, resp, err)
}

func CoinWalletValidateAddress(c *bm.Context) {
	p := new(ValidateAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.ValidateAddress(c, p)
	JF(c, resp, err)
}

func CoinWalletSendTo(c *bm.Context) {
	p := new(SendToRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.SendTo(c, p)
	JF(c, resp, err)
}

func CoinWalletSendToMany(c *bm.Context) {
	p := new(SendToManyRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.SendToMany(c, p)
	JF(c, resp, err)
}

func CoinWalletGetBalanceByAddress(c *bm.Context) {
	p := new(GetBalanceByAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetBalanceByAddress(c, p)
	JF(c, resp, err)
}

func CoinWalletGetBalance(c *bm.Context) {
	p := new(GetBalanceRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetBalance(c, p)
	JF(c, resp, err)
}

func CoinWalletAddMonitorAddress(c *bm.Context) {
	p := new(AddMonitorAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.AddMonitorAddress(c, p)
	JF(c, resp, err)
}

func CoinWalletRemoveMonitorAddress(c *bm.Context) {
	p := new(RemoveMonitorAddressRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.RemoveMonitorAddress(c, p)
	JF(c, resp, err)
}

func CoinWalletSetCallback(c *bm.Context) {
	p := new(SetCallbackRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.SetCallback(c, p)
	JF(c, resp, err)
}

func CoinWalletGetEstimateFee(c *bm.Context) {
	p := new(GetEstimateFeeRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetEstimateFee(c, p)
	JF(c, resp, err)
}

func CoinWalletGetMnemonicFingerprinting(c *bm.Context) {
	p := new(GetMnemonicFingerprintingRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.GetMnemonicFingerprinting(c, p)
	JF(c, resp, err)
}

func CoinWalletImportPrivateKey(c *bm.Context) {
	p := new(ImportPrivateKeyRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := CoinWalletSvc.ImportPrivateKey(c, p)
	JF(c, resp, err)
}

// RegisterCoinWalletBMServer Register the blademaster route
func RegisterCoinWalletBMServer(e *bm.Engine, server CoinWalletBMServer) {
	CoinWalletSvc = server
	e.GET(PathCoinWalletGetBlockInfo, CoinWalletGetBlockInfo)
	e.GET(PathCoinWalletGetTxinfo, CoinWalletGetTxinfo)
	e.GET(PathCoinWalletGetNewAddress, CoinWalletGetNewAddress)
	e.GET(PathCoinWalletValidateAddress, CoinWalletValidateAddress)
	e.GET(PathCoinWalletSendTo, CoinWalletSendTo)
	e.GET(PathCoinWalletSendToMany, CoinWalletSendToMany)
	e.GET(PathCoinWalletGetBalanceByAddress, CoinWalletGetBalanceByAddress)
	e.GET(PathCoinWalletGetBalance, CoinWalletGetBalance)
	e.GET(PathCoinWalletAddMonitorAddress, CoinWalletAddMonitorAddress)
	e.GET(PathCoinWalletRemoveMonitorAddress, CoinWalletRemoveMonitorAddress)
	e.GET(PathCoinWalletSetCallback, CoinWalletSetCallback)
	e.GET(PathCoinWalletGetEstimateFee, CoinWalletGetEstimateFee)
	e.GET(PathCoinWalletGetMnemonicFingerprinting, CoinWalletGetMnemonicFingerprinting)
	e.GET(PathCoinWalletImportPrivateKey, CoinWalletImportPrivateKey)
}
