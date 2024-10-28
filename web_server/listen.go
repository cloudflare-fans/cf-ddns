package web_server

import "github.com/cloudflare-fans/cf-ddns/sys_conf"

func (_this *CloudflareDDNSBackendWebServer) Listen() (err error) {
	err = _this.Engine.Run(sys_conf.GlobalSystemConf.Server.Listen)
	return
}
