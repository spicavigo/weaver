// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package chain

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)
var _ codegen.LatestVersion = codegen.Version[[0][17]struct{}]("You used 'weaver generate' codegen version 0.17.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/weavertest/internal/chain/A",
		Iface:       reflect.TypeOf((*A)(nil)).Elem(),
		Impl:        reflect.TypeOf(a{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return a_local_stub{impl: impl.(A), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return a_client_stub{stub: stub, propagateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/chain/A", Method: "Propagate"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return a_server_stub{impl: impl.(A), addLoad: addLoad}
		},
		RefData: "⟦d3d93f6e:wEaVeReDgE:github.com/ServiceWeaver/weaver/weavertest/internal/chain/A→github.com/ServiceWeaver/weaver/weavertest/internal/chain/B⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/weavertest/internal/chain/B",
		Iface:       reflect.TypeOf((*B)(nil)).Elem(),
		Impl:        reflect.TypeOf(b{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return b_local_stub{impl: impl.(B), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return b_client_stub{stub: stub, propagateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/chain/B", Method: "Propagate"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return b_server_stub{impl: impl.(B), addLoad: addLoad}
		},
		RefData: "⟦08d612ad:wEaVeReDgE:github.com/ServiceWeaver/weaver/weavertest/internal/chain/B→github.com/ServiceWeaver/weaver/weavertest/internal/chain/C⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/weavertest/internal/chain/C",
		Iface:       reflect.TypeOf((*C)(nil)).Elem(),
		Impl:        reflect.TypeOf(c{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return c_local_stub{impl: impl.(C), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return c_client_stub{stub: stub, propagateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/chain/C", Method: "Propagate"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return c_server_stub{impl: impl.(C), addLoad: addLoad}
		},
		RefData: "",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[A] = (*a)(nil)
var _ weaver.InstanceOf[B] = (*b)(nil)
var _ weaver.InstanceOf[C] = (*c)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*a)(nil)
var _ weaver.Unrouted = (*b)(nil)
var _ weaver.Unrouted = (*c)(nil)

// Local stub implementations.

type a_local_stub struct {
	impl   A
	tracer trace.Tracer
}

// Check that a_local_stub implements the A interface.
var _ A = (*a_local_stub)(nil)

func (s a_local_stub) Propagate(ctx context.Context, a0 int) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "chain.A.Propagate", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Propagate(ctx, a0)
}

type b_local_stub struct {
	impl   B
	tracer trace.Tracer
}

// Check that b_local_stub implements the B interface.
var _ B = (*b_local_stub)(nil)

func (s b_local_stub) Propagate(ctx context.Context, a0 int) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "chain.B.Propagate", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Propagate(ctx, a0)
}

type c_local_stub struct {
	impl   C
	tracer trace.Tracer
}

// Check that c_local_stub implements the C interface.
var _ C = (*c_local_stub)(nil)

func (s c_local_stub) Propagate(ctx context.Context, a0 int) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "chain.C.Propagate", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Propagate(ctx, a0)
}

// Client stub implementations.

type a_client_stub struct {
	stub             codegen.Stub
	propagateMetrics *codegen.MethodMetrics
}

// Check that a_client_stub implements the A interface.
var _ A = (*a_client_stub)(nil)

func (s a_client_stub) Propagate(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	start := time.Now()
	s.propagateMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "chain.A.Propagate", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.propagateMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.propagateMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.propagateMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.propagateMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type b_client_stub struct {
	stub             codegen.Stub
	propagateMetrics *codegen.MethodMetrics
}

// Check that b_client_stub implements the B interface.
var _ B = (*b_client_stub)(nil)

func (s b_client_stub) Propagate(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	start := time.Now()
	s.propagateMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "chain.B.Propagate", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.propagateMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.propagateMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.propagateMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.propagateMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type c_client_stub struct {
	stub             codegen.Stub
	propagateMetrics *codegen.MethodMetrics
}

// Check that c_client_stub implements the C interface.
var _ C = (*c_client_stub)(nil)

func (s c_client_stub) Propagate(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	start := time.Now()
	s.propagateMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "chain.C.Propagate", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.propagateMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.propagateMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.propagateMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.propagateMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Server stub implementations.

type a_server_stub struct {
	impl    A
	addLoad func(key uint64, load float64)
}

// Check that a_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*a_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s a_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Propagate":
		return s.propagate
	default:
		return nil
	}
}

func (s a_server_stub) propagate(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Propagate(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type b_server_stub struct {
	impl    B
	addLoad func(key uint64, load float64)
}

// Check that b_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*b_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s b_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Propagate":
		return s.propagate
	default:
		return nil
	}
}

func (s b_server_stub) propagate(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Propagate(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type c_server_stub struct {
	impl    C
	addLoad func(key uint64, load float64)
}

// Check that c_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*c_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s c_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Propagate":
		return s.propagate
	default:
		return nil
	}
}

func (s c_server_stub) propagate(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Propagate(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}
