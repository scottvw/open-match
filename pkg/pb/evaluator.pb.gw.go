// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/evaluator.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Evaluator_Evaluate_0(ctx context.Context, marshaler runtime.Marshaler, client EvaluatorClient, req *http.Request, pathParams map[string]string) (Evaluator_EvaluateClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.Evaluate(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq EvaluateRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Infof("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterEvaluatorHandlerFromEndpoint is same as RegisterEvaluatorHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEvaluatorHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEvaluatorHandler(ctx, mux, conn)
}

// RegisterEvaluatorHandler registers the http handlers for service Evaluator to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEvaluatorHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEvaluatorHandlerClient(ctx, mux, NewEvaluatorClient(conn))
}

// RegisterEvaluatorHandlerClient registers the http handlers for service Evaluator
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "EvaluatorClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "EvaluatorClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "EvaluatorClient" to call the correct interceptors.
func RegisterEvaluatorHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EvaluatorClient) error {

	mux.Handle("POST", pattern_Evaluator_Evaluate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Evaluator_Evaluate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Evaluator_Evaluate_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Evaluator_Evaluate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "evaluator", "matches"}, "evaluate", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Evaluator_Evaluate_0 = runtime.ForwardResponseStream
)
