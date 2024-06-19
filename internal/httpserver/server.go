package httpserver

import "github.com/gin-gonic/gin"

type HttpServer struct {
	host    string
	port    string
	ssl_crt string
	ssl_key string
	debug   bool
	ge      *gin.Engine
}

func NewHttpServer(host string, port string, ssl_crt string, ssl_key string, debug bool) *HttpServer {

	if !debug {

		gin.SetMode(gin.ReleaseMode)
	}

	return &HttpServer{
		host,
		port,
		ssl_crt,
		ssl_key,
		debug,
		gin.Default(),
	}
}
