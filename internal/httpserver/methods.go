package httpserver

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) GetRouter() *gin.Engine {

	return hs.ge
}

func (hs *HttpServer) RegisterGetRoute(path string, callback func(c *gin.Context)) {

	hs.ge.GET(path, callback)
}

func (hs *HttpServer) RegisterPostRoute(path string, callback func(c *gin.Context)) {

	hs.ge.POST(path, callback)
}

func (hs *HttpServer) Serve() error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	hs.ge.SetTrustedProxies(nil)

	go hs.ge.RunTLS(hs.host+":"+hs.port, pwd+hs.ssl_crt, pwd+hs.ssl_key)

	return nil
}
