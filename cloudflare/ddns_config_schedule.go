package cloudflare

import (
	"cf-ddns/util/duration_util"
	"errors"
	"time"
)

type DDNSScheduleInfo struct {
	DDNSName string
	Ticker   *time.Ticker
	Quitter  chan int
	Enabled  chan bool
}

type DDNSScheduleConfigsType map[string]DDNSScheduleInfo

var DDNSScheduleConfigs = DDNSScheduleConfigsType{}

func (_this *DDNSScheduleConfigsType) InitializeSchedule() error {
	if GlobalConfig.initialized == false {
		return errors.New("global config not initialized")
	}

	for _, target := range GlobalConfig.Targets {
		interval, err := duration_util.ParseNumTypedDuration(target.UpdateEvery)
		if err != nil {
			continue
		}

		ticker := time.NewTicker(interval)
		quitter := make(chan int)
		switcher := make(chan bool)

		go func() {
			for {
				select {
				case <-ticker.C:
					target.RunDDNS()
				case enabled := <-switcher:
					if enabled {
						ticker.Reset(interval)
					} else {
						ticker.Stop()
					}
				case <-quitter:
					ticker.Stop()
					return // pop the function stack
				}
			}
		}()
	}
	return nil
}
