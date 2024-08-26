/* Code generated by github.com/srdtrk/go-codegen, DO NOT EDIT. */
package schema

import (
	"context"
	"encoding/json"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Bucket is the client API for the QueryMsg_Bucket query message
	Bucket(ctx context.Context, req *QueryMsg_Bucket, opts ...grpc.CallOption) (*BucketResponse, error)
	// Object is the client API for the QueryMsg_Object query message
	Object(ctx context.Context, req *QueryMsg_Object, opts ...grpc.CallOption) (*ObjectResponse, error)
	// ObjectData is the client API for the QueryMsg_ObjectData query message
	ObjectData(ctx context.Context, req *QueryMsg_ObjectData, opts ...grpc.CallOption) (*string, error)
	// ObjectPins is the client API for the QueryMsg_ObjectPins query message
	ObjectPins(ctx context.Context, req *QueryMsg_ObjectPins, opts ...grpc.CallOption) (*ObjectPinsResponse, error)
	// Objects is the client API for the QueryMsg_Objects query message
	Objects(ctx context.Context, req *QueryMsg_Objects, opts ...grpc.CallOption) (*ObjectsResponse, error)
}

type queryClient struct {
	cc      *grpc.ClientConn
	address string
}

var _ QueryClient = (*queryClient)(nil)

// NewQueryClient creates a new QueryClient
func NewQueryClient(gRPCAddress, contractAddress string, opts ...grpc.DialOption) (QueryClient, error) {
	if len(opts) == 0 {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Create a connection to the gRPC server
	grpcConn, err := grpc.Dial(gRPCAddress, opts...)
	if err != nil {
		return nil, err
	}

	return &queryClient{
		address: contractAddress,
		cc:      grpcConn,
	}, nil
}

// Close closes the gRPC connection to the server
func (q *queryClient) Close() error {
	return q.cc.Close()
}

// queryContract is a helper function to query the contract with raw query data
func (q *queryClient) queryContract(ctx context.Context, rawQueryData []byte, opts ...grpc.CallOption) ([]byte, error) {
	in := &wasmtypes.QuerySmartContractStateRequest{
		Address:   q.address,
		QueryData: rawQueryData,
	}
	out := new(wasmtypes.QuerySmartContractStateResponse)
	err := q.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/SmartContractState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out.Data, nil
}

func (q *queryClient) Bucket(ctx context.Context, req *QueryMsg_Bucket, opts ...grpc.CallOption) (*BucketResponse, error) {
	rawQueryData, err := json.Marshal(map[string]any{"bucket": req})
	if err != nil {
		return nil, err
	}

	rawResponseData, err := q.queryContract(ctx, rawQueryData, opts...)
	if err != nil {
		return nil, err
	}

	var response BucketResponse
	if err := json.Unmarshal(rawResponseData, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (q *queryClient) Object(ctx context.Context, req *QueryMsg_Object, opts ...grpc.CallOption) (*ObjectResponse, error) {
	rawQueryData, err := json.Marshal(map[string]any{"object": req})
	if err != nil {
		return nil, err
	}

	rawResponseData, err := q.queryContract(ctx, rawQueryData, opts...)
	if err != nil {
		return nil, err
	}

	var response ObjectResponse
	if err := json.Unmarshal(rawResponseData, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (q *queryClient) ObjectData(ctx context.Context, req *QueryMsg_ObjectData, opts ...grpc.CallOption) (*string, error) {
	rawQueryData, err := json.Marshal(map[string]any{"object_data": req})
	if err != nil {
		return nil, err
	}

	rawResponseData, err := q.queryContract(ctx, rawQueryData, opts...)
	if err != nil {
		return nil, err
	}

	var response string
	if err := json.Unmarshal(rawResponseData, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (q *queryClient) ObjectPins(ctx context.Context, req *QueryMsg_ObjectPins, opts ...grpc.CallOption) (*ObjectPinsResponse, error) {
	rawQueryData, err := json.Marshal(map[string]any{"object_pins": req})
	if err != nil {
		return nil, err
	}

	rawResponseData, err := q.queryContract(ctx, rawQueryData, opts...)
	if err != nil {
		return nil, err
	}

	var response ObjectPinsResponse
	if err := json.Unmarshal(rawResponseData, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (q *queryClient) Objects(ctx context.Context, req *QueryMsg_Objects, opts ...grpc.CallOption) (*ObjectsResponse, error) {
	rawQueryData, err := json.Marshal(map[string]any{"objects": req})
	if err != nil {
		return nil, err
	}

	rawResponseData, err := q.queryContract(ctx, rawQueryData, opts...)
	if err != nil {
		return nil, err
	}

	var response ObjectsResponse
	if err := json.Unmarshal(rawResponseData, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
