package web_server

import "github.com/gin-gonic/gin"

type CloudflareDDNSBackendWebServer struct {
	Engine *gin.Engine
}

func Init() *CloudflareDDNSBackendWebServer {
	return &CloudflareDDNSBackendWebServer{Engine: gin.Default()}
}
