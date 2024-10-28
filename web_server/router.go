package web_server

func (_this *CloudflareDDNSBackendWebServer) RegisterRouter() {
	v1 := _this.Engine.Group("/api/v1/")
	{
		v1.POST("/ddns-tasks/all/-run")
	}
}
