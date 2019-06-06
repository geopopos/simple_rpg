// TODO pack into middleware util
package logging

import (
	server "github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

var Formatter = &log.JSONFormatter{}

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.SetFormatter(Formatter)
		log.Infof("Server Request: %s", req.Method())
		return fn(ctx, req, rsp)
	}
}
