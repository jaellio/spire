// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/api/node (interfaces: NodeClient,Node_FetchSVIDClient,NodeServer,Node_FetchSVIDServer)

package mock_node

import (
	gomock "github.com/golang/mock/gomock"
	node "github.com/spiffe/spire/proto/api/node"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockNodeClient is a mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return _m.recorder
}

// FetchBaseSVID mocks base method
func (_m *MockNodeClient) FetchBaseSVID(_param0 context.Context, _param1 *node.FetchBaseSVIDRequest, _param2 ...grpc.CallOption) (*node.FetchBaseSVIDResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "FetchBaseSVID", _s...)
	ret0, _ := ret[0].(*node.FetchBaseSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBaseSVID indicates an expected call of FetchBaseSVID
func (_mr *MockNodeClientMockRecorder) FetchBaseSVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchBaseSVID", reflect.TypeOf((*MockNodeClient)(nil).FetchBaseSVID), _s...)
}

// FetchCPBundle mocks base method
func (_m *MockNodeClient) FetchCPBundle(_param0 context.Context, _param1 *node.FetchCPBundleRequest, _param2 ...grpc.CallOption) (*node.FetchCPBundleResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "FetchCPBundle", _s...)
	ret0, _ := ret[0].(*node.FetchCPBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCPBundle indicates an expected call of FetchCPBundle
func (_mr *MockNodeClientMockRecorder) FetchCPBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchCPBundle", reflect.TypeOf((*MockNodeClient)(nil).FetchCPBundle), _s...)
}

// FetchFederatedBundle mocks base method
func (_m *MockNodeClient) FetchFederatedBundle(_param0 context.Context, _param1 *node.FetchFederatedBundleRequest, _param2 ...grpc.CallOption) (*node.FetchFederatedBundleResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "FetchFederatedBundle", _s...)
	ret0, _ := ret[0].(*node.FetchFederatedBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (_mr *MockNodeClientMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockNodeClient)(nil).FetchFederatedBundle), _s...)
}

// FetchSVID mocks base method
func (_m *MockNodeClient) FetchSVID(_param0 context.Context, _param1 ...grpc.CallOption) (node.Node_FetchSVIDClient, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "FetchSVID", _s...)
	ret0, _ := ret[0].(node.Node_FetchSVIDClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSVID indicates an expected call of FetchSVID
func (_mr *MockNodeClientMockRecorder) FetchSVID(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchSVID", reflect.TypeOf((*MockNodeClient)(nil).FetchSVID), _s...)
}

// MockNode_FetchSVIDClient is a mock of Node_FetchSVIDClient interface
type MockNode_FetchSVIDClient struct {
	ctrl     *gomock.Controller
	recorder *MockNode_FetchSVIDClientMockRecorder
}

// MockNode_FetchSVIDClientMockRecorder is the mock recorder for MockNode_FetchSVIDClient
type MockNode_FetchSVIDClientMockRecorder struct {
	mock *MockNode_FetchSVIDClient
}

// NewMockNode_FetchSVIDClient creates a new mock instance
func NewMockNode_FetchSVIDClient(ctrl *gomock.Controller) *MockNode_FetchSVIDClient {
	mock := &MockNode_FetchSVIDClient{ctrl: ctrl}
	mock.recorder = &MockNode_FetchSVIDClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNode_FetchSVIDClient) EXPECT() *MockNode_FetchSVIDClientMockRecorder {
	return _m.recorder
}

// CloseSend mocks base method
func (_m *MockNode_FetchSVIDClient) CloseSend() error {
	ret := _m.ctrl.Call(_m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (_mr *MockNode_FetchSVIDClientMockRecorder) CloseSend() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CloseSend", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).CloseSend))
}

// Context mocks base method
func (_m *MockNode_FetchSVIDClient) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockNode_FetchSVIDClientMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).Context))
}

// Header mocks base method
func (_m *MockNode_FetchSVIDClient) Header() (metadata.MD, error) {
	ret := _m.ctrl.Call(_m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (_mr *MockNode_FetchSVIDClientMockRecorder) Header() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Header", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).Header))
}

// Recv mocks base method
func (_m *MockNode_FetchSVIDClient) Recv() (*node.FetchSVIDResponse, error) {
	ret := _m.ctrl.Call(_m, "Recv")
	ret0, _ := ret[0].(*node.FetchSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (_mr *MockNode_FetchSVIDClientMockRecorder) Recv() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Recv", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).Recv))
}

// RecvMsg mocks base method
func (_m *MockNode_FetchSVIDClient) RecvMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "RecvMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (_mr *MockNode_FetchSVIDClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (_m *MockNode_FetchSVIDClient) Send(_param0 *node.FetchSVIDRequest) error {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (_mr *MockNode_FetchSVIDClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Send", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (_m *MockNode_FetchSVIDClient) SendMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SendMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (_mr *MockNode_FetchSVIDClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendMsg", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (_m *MockNode_FetchSVIDClient) Trailer() metadata.MD {
	ret := _m.ctrl.Call(_m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (_mr *MockNode_FetchSVIDClientMockRecorder) Trailer() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Trailer", reflect.TypeOf((*MockNode_FetchSVIDClient)(nil).Trailer))
}

// MockNodeServer is a mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return _m.recorder
}

// FetchBaseSVID mocks base method
func (_m *MockNodeServer) FetchBaseSVID(_param0 context.Context, _param1 *node.FetchBaseSVIDRequest) (*node.FetchBaseSVIDResponse, error) {
	ret := _m.ctrl.Call(_m, "FetchBaseSVID", _param0, _param1)
	ret0, _ := ret[0].(*node.FetchBaseSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBaseSVID indicates an expected call of FetchBaseSVID
func (_mr *MockNodeServerMockRecorder) FetchBaseSVID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchBaseSVID", reflect.TypeOf((*MockNodeServer)(nil).FetchBaseSVID), arg0, arg1)
}

// FetchCPBundle mocks base method
func (_m *MockNodeServer) FetchCPBundle(_param0 context.Context, _param1 *node.FetchCPBundleRequest) (*node.FetchCPBundleResponse, error) {
	ret := _m.ctrl.Call(_m, "FetchCPBundle", _param0, _param1)
	ret0, _ := ret[0].(*node.FetchCPBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCPBundle indicates an expected call of FetchCPBundle
func (_mr *MockNodeServerMockRecorder) FetchCPBundle(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchCPBundle", reflect.TypeOf((*MockNodeServer)(nil).FetchCPBundle), arg0, arg1)
}

// FetchFederatedBundle mocks base method
func (_m *MockNodeServer) FetchFederatedBundle(_param0 context.Context, _param1 *node.FetchFederatedBundleRequest) (*node.FetchFederatedBundleResponse, error) {
	ret := _m.ctrl.Call(_m, "FetchFederatedBundle", _param0, _param1)
	ret0, _ := ret[0].(*node.FetchFederatedBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (_mr *MockNodeServerMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockNodeServer)(nil).FetchFederatedBundle), arg0, arg1)
}

// FetchSVID mocks base method
func (_m *MockNodeServer) FetchSVID(_param0 node.Node_FetchSVIDServer) error {
	ret := _m.ctrl.Call(_m, "FetchSVID", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchSVID indicates an expected call of FetchSVID
func (_mr *MockNodeServerMockRecorder) FetchSVID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchSVID", reflect.TypeOf((*MockNodeServer)(nil).FetchSVID), arg0)
}

// MockNode_FetchSVIDServer is a mock of Node_FetchSVIDServer interface
type MockNode_FetchSVIDServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_FetchSVIDServerMockRecorder
}

// MockNode_FetchSVIDServerMockRecorder is the mock recorder for MockNode_FetchSVIDServer
type MockNode_FetchSVIDServerMockRecorder struct {
	mock *MockNode_FetchSVIDServer
}

// NewMockNode_FetchSVIDServer creates a new mock instance
func NewMockNode_FetchSVIDServer(ctrl *gomock.Controller) *MockNode_FetchSVIDServer {
	mock := &MockNode_FetchSVIDServer{ctrl: ctrl}
	mock.recorder = &MockNode_FetchSVIDServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNode_FetchSVIDServer) EXPECT() *MockNode_FetchSVIDServerMockRecorder {
	return _m.recorder
}

// Context mocks base method
func (_m *MockNode_FetchSVIDServer) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockNode_FetchSVIDServerMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).Context))
}

// Recv mocks base method
func (_m *MockNode_FetchSVIDServer) Recv() (*node.FetchSVIDRequest, error) {
	ret := _m.ctrl.Call(_m, "Recv")
	ret0, _ := ret[0].(*node.FetchSVIDRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (_mr *MockNode_FetchSVIDServerMockRecorder) Recv() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Recv", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).Recv))
}

// RecvMsg mocks base method
func (_m *MockNode_FetchSVIDServer) RecvMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "RecvMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (_mr *MockNode_FetchSVIDServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (_m *MockNode_FetchSVIDServer) Send(_param0 *node.FetchSVIDResponse) error {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (_mr *MockNode_FetchSVIDServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Send", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (_m *MockNode_FetchSVIDServer) SendHeader(_param0 metadata.MD) error {
	ret := _m.ctrl.Call(_m, "SendHeader", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (_mr *MockNode_FetchSVIDServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendHeader", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (_m *MockNode_FetchSVIDServer) SendMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SendMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (_mr *MockNode_FetchSVIDServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendMsg", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (_m *MockNode_FetchSVIDServer) SetHeader(_param0 metadata.MD) error {
	ret := _m.ctrl.Call(_m, "SetHeader", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (_mr *MockNode_FetchSVIDServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetHeader", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (_m *MockNode_FetchSVIDServer) SetTrailer(_param0 metadata.MD) {
	_m.ctrl.Call(_m, "SetTrailer", _param0)
}

// SetTrailer indicates an expected call of SetTrailer
func (_mr *MockNode_FetchSVIDServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_FetchSVIDServer)(nil).SetTrailer), arg0)
}
