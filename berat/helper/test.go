package helper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockInterfaceMethod interface {
	On(methodName string, arguments ...interface{}) *mock.Call
}

type TestFrameworkReq struct {
	Cases    []Cases
	T        *testing.T
	TestFunc func(interface{}) []interface{}
	BaseReq  interface{}
}

type Cases struct {
	MockFunc []MockFunc
	Want     []interface{}
	PatchReq func(interface{}) interface{}
	Name     string
}

type MockFunc struct {
	TestForErr    TestForErr
	MockInterface MockInterfaceMethod
	AddToMock     bool
	MockFuncName  string
	MockReq       []interface{}
	MockRes       []interface{}
}

type TestForErr struct {
	Test        bool
	ErrPosition int
}

func TestFramework(req TestFrameworkReq) {
	mockFuncAll := []func(){}

	for _, mockCase := range req.Cases {
		tempMock := []func(){}
		for _, mockFunc := range mockCase.MockFunc {
			fn := func() {
				mockFunc.MockInterface.On(mockFunc.MockFuncName, mockFunc.MockReq...).Return(mockFunc.MockRes...).Once()
			}

			if mockFunc.TestForErr.Test {
				res := make([]interface{}, len(mockFunc.MockRes))
				for i, data := range mockFunc.MockRes {
					res[i] = data
				}
				res[mockFunc.TestForErr.ErrPosition] = errors.New(fmt.Sprintf("err %s", mockFunc.MockFuncName))
				fnErr := func() {
					mockFunc.MockInterface.On(mockFunc.MockFuncName, mockFunc.MockReq...).Return(res...).Once()
				}

				tempForErr := make([]func(), len(mockFuncAll)+1)
				for i, data := range mockFuncAll {
					tempForErr[i] = data
				}
				tempForErr[len(mockFuncAll)] = fnErr
				RunTest(RunTestReq{
					T:           req.T,
					CaseName:    fmt.Sprintf("err %s", mockFunc.MockFuncName),
					MockFuncAll: tempForErr,
					BaseReq:     req.BaseReq,
					PatchReq:    mockCase.PatchReq,
					TestFunc:    req.TestFunc,
					Want:        []interface{}{errors.New(fmt.Sprintf("err %s", mockFunc.MockFuncName))},
					TestForErr: TestForErr{
						Test:        true,
						ErrPosition: mockFunc.TestForErr.ErrPosition,
					},
				})
			}

			if mockFunc.AddToMock {
				mockFuncAll = append(mockFuncAll, fn)
				continue
			}
			tempMock = append(tempMock, fn)
		}

		RunTest(RunTestReq{
			T:           req.T,
			CaseName:    mockCase.Name,
			MockFuncAll: append(mockFuncAll, tempMock...),
			BaseReq:     req.BaseReq,
			PatchReq:    mockCase.PatchReq,
			TestFunc:    req.TestFunc,
			Want:        mockCase.Want,
		})
	}
}

type RunTestReq struct {
	T           *testing.T
	CaseName    string
	MockFuncAll []func()
	BaseReq     interface{}
	PatchReq    func(interface{}) interface{}
	TestFunc    func(interface{}) []interface{}
	Want        []interface{}
	TestForErr  TestForErr
}

func RunTest(req RunTestReq) {
	req.T.Run(req.CaseName, func(*testing.T) {
		for _, mockFn := range req.MockFuncAll {
			mockFn()
		}

		reqTest := req.BaseReq
		if req.PatchReq != nil {
			reqTest = req.PatchReq(reqTest)
		}

		gotFromFunc := req.TestFunc(reqTest)

		if req.TestForErr.Test {
			if !reflect.DeepEqual(gotFromFunc[req.TestForErr.ErrPosition], req.Want[0]) {
				req.T.Errorf("Test %s failed, got: %+v, want: %+v", req.CaseName, gotFromFunc[req.TestForErr.ErrPosition], req.Want[0])
			}
			return
		}

		if !reflect.DeepEqual(gotFromFunc, req.Want) {
			req.T.Errorf("Test %s failed, got: %+v, want: %+v", req.CaseName, gotFromFunc, req.Want)
		}
	})
}
