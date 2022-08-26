// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PeerCount provides a mock function with given fields: ctx
func (_m *EthClient) PeerCount(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *EthClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	var r0 ethereum.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncProgress provides a mock function with given fields: ctx
func (_m *EthClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	ret := _m.Called(ctx)

	var r0 *ethereum.SyncProgress
	if rf, ok := ret.Get(0).(func(context.Context) *ethereum.SyncProgress); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethereum.SyncProgress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEthClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewEthClient creates a new instance of EthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEthClient(t mockConstructorTestingTNewEthClient) *EthClient {
	mock := &EthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}