package pmsbase

import (
	"quant/hqmodule/base"
)

// IStrategy define the function which should be implemented by strategy
type IStrategy interface {
	// implement by strategy
	SInit()
	CalcSignal(*hqbase.Marketdata) int
	GeneratePortfolio(int, *hqbase.Marketdata)
	// implement by sbase
	Stop()
	Start()
	GetID() string
	GetAccount() string
	GetTradeStatus() string
	GetSecurityID() string
}

// SbaseConfig is config of Sbase struct
type SbaseConfig struct {
	SubQuoteCodes string
	AdapterName   string
	AccountID     string
	CombiNo       string
}

// StrategyID is used for new strategy response
type StrategyID struct {
	Ret int
	Msg string
	ID  string
}

// StrategyTemp is used for getStrategyTemp
type StrategyTemp struct {
	StrategyName   string
	Author         string
	MDD1           float64
	MDD1SD         int
	MDD1ED         int
	MDD2           float64
	MDD2SD         int
	MDD2ED         int
	AnnRet         float64
	Vol            float64
	Calmar         float64
	SR             float64
	Reamrk         string
	CreateDateTime string
}

// StrategyRunningInfo is the information of running strategies
type StrategyRunningInfo struct {
	StrategyID   string
	Account      string
	StrategyName string
	SecurityID   string
	RunStatus    string // RUNNING/STOP
	TradeStatus  string // TRADING/-
}
