// This file was generated by counterfeiter
package extensionfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/pivotal-cf/go-pivnet/extension"
)

type FakeClient struct {
	MakeRequestStub        func(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, error)
	makeRequestMutex       sync.RWMutex
	makeRequestArgsForCall []struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}
	makeRequestReturns struct {
		result1 *http.Response
		result2 error
	}
	CreateRequestStub        func(method string, url string, body io.Reader) (*http.Request, error)
	createRequestMutex       sync.RWMutex
	createRequestArgsForCall []struct {
		method string
		url    string
		body   io.Reader
	}
	createRequestReturns struct {
		result1 *http.Request
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) MakeRequest(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, error) {
	fake.makeRequestMutex.Lock()
	fake.makeRequestArgsForCall = append(fake.makeRequestArgsForCall, struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}{method, url, expectedResponseCode, body, data})
	fake.recordInvocation("MakeRequest", []interface{}{method, url, expectedResponseCode, body, data})
	fake.makeRequestMutex.Unlock()
	if fake.MakeRequestStub != nil {
		return fake.MakeRequestStub(method, url, expectedResponseCode, body, data)
	} else {
		return fake.makeRequestReturns.result1, fake.makeRequestReturns.result2
	}
}

func (fake *FakeClient) MakeRequestCallCount() int {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return len(fake.makeRequestArgsForCall)
}

func (fake *FakeClient) MakeRequestArgsForCall(i int) (string, string, int, io.Reader, interface{}) {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return fake.makeRequestArgsForCall[i].method, fake.makeRequestArgsForCall[i].url, fake.makeRequestArgsForCall[i].expectedResponseCode, fake.makeRequestArgsForCall[i].body, fake.makeRequestArgsForCall[i].data
}

func (fake *FakeClient) MakeRequestReturns(result1 *http.Response, result2 error) {
	fake.MakeRequestStub = nil
	fake.makeRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateRequest(method string, url string, body io.Reader) (*http.Request, error) {
	fake.createRequestMutex.Lock()
	fake.createRequestArgsForCall = append(fake.createRequestArgsForCall, struct {
		method string
		url    string
		body   io.Reader
	}{method, url, body})
	fake.recordInvocation("CreateRequest", []interface{}{method, url, body})
	fake.createRequestMutex.Unlock()
	if fake.CreateRequestStub != nil {
		return fake.CreateRequestStub(method, url, body)
	} else {
		return fake.createRequestReturns.result1, fake.createRequestReturns.result2
	}
}

func (fake *FakeClient) CreateRequestCallCount() int {
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	return len(fake.createRequestArgsForCall)
}

func (fake *FakeClient) CreateRequestArgsForCall(i int) (string, string, io.Reader) {
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	return fake.createRequestArgsForCall[i].method, fake.createRequestArgsForCall[i].url, fake.createRequestArgsForCall[i].body
}

func (fake *FakeClient) CreateRequestReturns(result1 *http.Request, result2 error) {
	fake.CreateRequestStub = nil
	fake.createRequestReturns = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ extension.Client = new(FakeClient)