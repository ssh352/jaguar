package helper

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/hqmodule/base"
	redis "util/redis"
)

var (
	conf        = goini.SetConfig(QuantConfigFile)
	redisconfig = map[string]string{
		"host":         conf.GetValue("redis", "redishost"),
		"database":     conf.GetValue("redis", "database"),
		"password":     conf.GetValue("redis", "password"),
		"maxOpenConns": conf.GetValue("redis", "maxOpenConns"),
		"maxIdleConns": conf.GetValue("redis", "maxIdleConns"),
	}
	redisPool *redis.ConnPool
)

func init() {
	redisPool = redis.InitRedis(redisconfig)
	_, err := redisPool.Do("PING")
	if err != nil {
		log.Error("redis error")
		panic(err)
	}
}

// GetQuote get market data from redis
func GetQuote(code string) (hqbase.Marketdata, error) {

	tmp, err := redisPool.GetHashMapKey(RedisKey, code)

	if err != nil {
		return hqbase.Marketdata{}, &Error{fmt.Sprintf("Quant helper GetQuote error.\"%s\" ", err.Error())}
	}
	var mkdat hqbase.Marketdata
	err = msgpack.Unmarshal(tmp.([]byte), &mkdat)
	if err != nil {
		return hqbase.Marketdata{}, &Error{fmt.Sprintf("Quant helper GetQuote error. \"%s\"", err.Error())}
	}
	return mkdat, nil
}

// GetQuotes get market datas from redis
func GetQuotes(codes []string) (map[string]hqbase.Marketdata, error) {

	quotemap := make(map[string]hqbase.Marketdata)
	tmp, err := redisPool.GetMHashMapString(RedisKey, codes)
	if err != nil {
		return quotemap, &Error{fmt.Sprintf("Quant helper GetQuotes error.\"%s\" ", err.Error())}
	}
	for i, quote := range tmp {
		var mkdat hqbase.Marketdata
		msgpack.Unmarshal([]byte(quote), &mkdat)
		quotemap[codes[i]] = mkdat
	}
	return quotemap, nil
}
