package main

import (
	log "github.com/thinkboy/log4go"
	"quant/helper"
	"quant/rmsmodule/base"
	"time"
	"util/csp"
)

var (
	riskrules      map[string]rmsbase.RiskRule
	riskfuncrouter map[string]func(*csp.Request) csp.Response
)

func init() {
	logfiles := make(map[string]string)
	logfiles["ERROR"] = "rmsmodule_err%s.log" + time.Now().Format("2006-01-02")
	logfiles["INFO"] = "rmsmodule_info%s.log" + time.Now().Format("2006-01-02")
	log.SetLogFiles(logfiles)
	log.LoadConfiguration(helper.QuantLogConfigFile)
	// riskrules = make([]rmsbase.RiskRule, 1)
	riskfuncrouter = make(map[string]func(*csp.Request) csp.Response, 10)
	riskrules = make(map[string]rmsbase.RiskRule, 10)
}

func main() {
	admin := riskAdmin{}
	admin.init()
	riskfuncrouter["getTradePct"] = admin.getTradePct
	riskfuncrouter["getEntrustAmountInfo"] = admin.getEntrustAmountInfo
	riskfuncrouter["getRiskRules"] = admin.getRiskRules
	riskfuncrouter["modifyRiskRules"] = admin.modifyRiskRules
	riskfuncrouter["delRiskRules"] = admin.delRiskRules
	riskfuncrouter["addRiskRules"] = admin.addRiskRules
	newService()
	wc := make(chan int)
	<-wc
}
