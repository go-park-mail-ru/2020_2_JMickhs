package grpcPackage

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"google.golang.org/grpc"
)

func GetInterceptor(log *logger.CustomLogger) func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Tracef("call=%v req=%#v reply=%#v time=%v err=%v",
			method, req, reply, time.Since(start), err)
		return err
	}
}
