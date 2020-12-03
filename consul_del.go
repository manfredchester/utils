package main

import (
	"msp-git.connext.com.cn/module/pub/zhlog"

	"github.com/hashicorp/consul/api"
)

func DelConsulInfo(key, consulAddr string) {
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error(zhlog.UUID(8), "删除Consul配置发生错误:", e.(error))
		}
	}()
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	zhlog.Assert(err)
	_, err = consulClient.KV().Delete(key, nil)
	zhlog.Assert(err)
}
