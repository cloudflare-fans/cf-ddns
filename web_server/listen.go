package web_server

func (_this *CloudflareDDNSBackendWebServer) Listen() (err error) {
	err = _this.Engine.Run()
	return
}
