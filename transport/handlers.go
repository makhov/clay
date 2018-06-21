package transport

import (
	"net/http"

	"github.com/utrack/clay/transport/swagger"
	"google.golang.org/grpc"
)

// Service is a registerable collection of endpoints.
// These functions should be autogenerated by protoc-gen-goclay.
type Service interface {
	GetDescription() ServiceDesc
}

// ServiceDesc is a description of an endpoints' collection.
// These functions should be autogenerated by protoc-gen-goclay.
type ServiceDesc interface {
	RegisterGRPC(*grpc.Server)
	RegisterHTTP(Router) error
	SwaggerDef(options ...swagger.Option) []byte
}

// Router routes HTTP requests around.
type Router interface {
	Handle(pattern string, h http.Handler)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
