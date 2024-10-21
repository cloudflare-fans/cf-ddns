package main

import (
	"cf-ddns/cloudflare"
	"time"
)

func RunUpdate(cf cloudflare.Cloudflare) (err error) {
	err = cf.DetectCurrentIP()
	if err != nil {
		return
	}
	err = cf.DetectConfiguredRecord()
	if err != nil {
		return
	}
	if cf.ShouldUpdate() {
		// 重复五次提交
		for i := 0; i < 5; i++ {
			err = cf.UpdateDNSRecord(false)
			if err == nil {
				break
			} else {
				time.Sleep(5 * time.Second)
			}
		}
	}
	return
}
