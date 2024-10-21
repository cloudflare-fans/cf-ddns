package cloudflare

import (
	"log"
	"time"
)

func (_this *cloudflareDDNSClient) RunDDNS() (err error) {
	err = _this.detectCurrentIP()
	if err != nil {
		return
	}
	err = _this.detectConfiguredRecord()
	if err != nil {
		return
	}
	if _this.shouldUpdate() {
		// 重复五次提交
		for i := 0; i < 5; i++ {
			log.Printf("Updating DNS record (trying %v)\n", i+1)
			err = _this.updateDNSRecord(false)
			if err == nil {
				break
			} else {
				log.Printf("Failed to update DNS record (trying %v): %s\n", i+1, err)
				time.Sleep(5 * time.Second)
			}
		}
	}
	return
}
