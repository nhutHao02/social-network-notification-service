package api

import "github.com/nhutHao02/social-network-notification-service/internal/api/http"

type Server struct {
	// http server
	HTTPServer *http.HTTPServer
	// grpc server

}

func NewSerVer(httpServer *http.HTTPServer) *Server {
	return &Server{HTTPServer: httpServer}
}
