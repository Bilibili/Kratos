// Code generated by protoc-gen-go-http. DO NOT EDIT.

package helloworld

import (
	context "context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	transport "github.com/go-kratos/kratos/v2/transport/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context.Context
const _ = transport.SupportPackageIsVersion1

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterHTTPServer(s transport.ServiceRegistrar, srv GreeterHTTPServer) {
	s.RegisterService(&_HTTP_Greeter_serviceDesc, srv)
}

func _HTTP_Greeter_SayHello(srv interface{}, ctx context.Context, dec func(interface{}) error, req *http.Request) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	vars := transport.PathParams(req)
	name, ok := vars["name"]
	if !ok {
		return nil, errors.InvalidArgument("Errors_InvalidArgument", "missing parameter: name")
	}
	in.Name = name
	reply, err := srv.(GreeterServer).SayHello(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

var _HTTP_Greeter_serviceDesc = transport.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterHTTPServer)(nil),
	Methods: []transport.MethodDesc{

		{
			Path:    "/helloworld/{name}",
			Method:  "GET",
			Handler: _HTTP_Greeter_SayHello,
		},
	},
	Metadata: "helloworld.proto",
}