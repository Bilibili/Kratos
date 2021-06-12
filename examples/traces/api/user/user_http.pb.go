// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	transport "github.com/go-kratos/kratos/v2/transport"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(middleware.Middleware)
var _ = new(transport.Transporter)
var _ = binding.BindVars
var _ = mux.NewRouter

const _ = http.SupportPackageIsVersion1

type UserHandler interface {
	GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHandler) {
	r := s.Route("/")

	r.GET("/v1/user/get/message/{count}", func(ctx http.Context) error {
		var in GetMyMessagesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}

		if err := binding.BindVars(ctx.Vars(), &in); err != nil {
			return err
		}

		transport.SetOperation(ctx, "/api.user.v1.User/GetMyMessages")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyMessages(ctx, req.(*GetMyMessagesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyMessagesReply)
		return ctx.Result(200, reply)
	})

}

type UserHTTPClient interface {
	GetMyMessages(ctx context.Context, req *GetMyMessagesRequest, opts ...http.CallOption) (rsp *GetMyMessagesReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetMyMessages(ctx context.Context, in *GetMyMessagesRequest, opts ...http.CallOption) (*GetMyMessagesReply, error) {
	var out GetMyMessagesReply
	path := binding.EncodeVars("/v1/user/get/message/{count}", in, true)
	opts = append(opts, http.Operation("/api.user.v1.User/GetMyMessages"))

	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)

	return &out, err
}
