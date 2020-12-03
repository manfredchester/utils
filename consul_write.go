package main

import (
	"fmt"
	"path"

	"github.com/hashicorp/consul/api"
	"msp-git.connext.com.cn/module/pub/zhlog"
)

func WriteConsulInfo(key, val, consulAddr string) {
	traceID := zhlog.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error(traceID, "写入Consul配置发生错误:", e.(error))
		}
	}()
	if path.IsAbs(key) {
		zhlog.Assert(fmt.Errorf("%s", "key应为相对路径"))
	}
	value := []byte(val)
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	zhlog.Assert(err)
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, nil)
	zhlog.Assert(err)
}
