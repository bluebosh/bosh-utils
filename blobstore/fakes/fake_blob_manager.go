// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/bluebosh/bosh-utils/blobstore"
	"github.com/bluebosh/bosh-utils/crypto"
	"github.com/bluebosh/bosh-utils/system"
)

type FakeBlobManagerInterface struct {
	FetchStub        func(blobID string) (system.File, error, int)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		blobID string
	}
	fetchReturns struct {
		result1 system.File
		result2 error
		result3 int
	}
	WriteStub        func(blobID string, reader io.Reader) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		blobID string
		reader io.Reader
	}
	writeReturns struct {
		result1 error
	}
	GetPathStub        func(blobID string, digest crypto.Digest) (string, error)
	getPathMutex       sync.RWMutex
	getPathArgsForCall []struct {
		blobID string
		digest crypto.Digest
	}
	getPathReturns struct {
		result1 string
		result2 error
	}
	DeleteStub        func(blobID string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		blobID string
	}
	deleteReturns struct {
		result1 error
	}
	BlobExistsStub        func(blobID string) bool
	blobExistsMutex       sync.RWMutex
	blobExistsArgsForCall []struct {
		blobID string
	}
	blobExistsReturns struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobManagerInterface) Fetch(blobID string) (system.File, error, int) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		blobID string
	}{blobID})
	fake.recordInvocation("Fetch", []interface{}{blobID})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(blobID)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2, fake.fetchReturns.result3
	}
}

func (fake *FakeBlobManagerInterface) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeBlobManagerInterface) FetchArgsForCall(i int) string {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].blobID
}

func (fake *FakeBlobManagerInterface) FetchReturns(result1 system.File, result2 error, result3 int) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 system.File
		result2 error
		result3 int
	}{result1, result2, result3}
}

func (fake *FakeBlobManagerInterface) Write(blobID string, reader io.Reader) error {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		blobID string
		reader io.Reader
	}{blobID, reader})
	fake.recordInvocation("Write", []interface{}{blobID, reader})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(blobID, reader)
	} else {
		return fake.writeReturns.result1
	}
}

func (fake *FakeBlobManagerInterface) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeBlobManagerInterface) WriteArgsForCall(i int) (string, io.Reader) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].blobID, fake.writeArgsForCall[i].reader
}

func (fake *FakeBlobManagerInterface) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobManagerInterface) GetPath(blobID string, digest crypto.Digest) (string, error) {
	fake.getPathMutex.Lock()
	fake.getPathArgsForCall = append(fake.getPathArgsForCall, struct {
		blobID string
		digest crypto.Digest
	}{blobID, digest})
	fake.recordInvocation("GetPath", []interface{}{blobID, digest})
	fake.getPathMutex.Unlock()
	if fake.GetPathStub != nil {
		return fake.GetPathStub(blobID, digest)
	} else {
		return fake.getPathReturns.result1, fake.getPathReturns.result2
	}
}

func (fake *FakeBlobManagerInterface) GetPathCallCount() int {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	return len(fake.getPathArgsForCall)
}

func (fake *FakeBlobManagerInterface) GetPathArgsForCall(i int) (string, crypto.Digest) {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	return fake.getPathArgsForCall[i].blobID, fake.getPathArgsForCall[i].digest
}

func (fake *FakeBlobManagerInterface) GetPathReturns(result1 string, result2 error) {
	fake.GetPathStub = nil
	fake.getPathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobManagerInterface) Delete(blobID string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		blobID string
	}{blobID})
	fake.recordInvocation("Delete", []interface{}{blobID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(blobID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeBlobManagerInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeBlobManagerInterface) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].blobID
}

func (fake *FakeBlobManagerInterface) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobManagerInterface) BlobExists(blobID string) bool {
	fake.blobExistsMutex.Lock()
	fake.blobExistsArgsForCall = append(fake.blobExistsArgsForCall, struct {
		blobID string
	}{blobID})
	fake.recordInvocation("BlobExists", []interface{}{blobID})
	fake.blobExistsMutex.Unlock()
	if fake.BlobExistsStub != nil {
		return fake.BlobExistsStub(blobID)
	} else {
		return fake.blobExistsReturns.result1
	}
}

func (fake *FakeBlobManagerInterface) BlobExistsCallCount() int {
	fake.blobExistsMutex.RLock()
	defer fake.blobExistsMutex.RUnlock()
	return len(fake.blobExistsArgsForCall)
}

func (fake *FakeBlobManagerInterface) BlobExistsArgsForCall(i int) string {
	fake.blobExistsMutex.RLock()
	defer fake.blobExistsMutex.RUnlock()
	return fake.blobExistsArgsForCall[i].blobID
}

func (fake *FakeBlobManagerInterface) BlobExistsReturns(result1 bool) {
	fake.BlobExistsStub = nil
	fake.blobExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBlobManagerInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.blobExistsMutex.RLock()
	defer fake.blobExistsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBlobManagerInterface) recordInvocation(key string, args []interface{}) {
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

var _ blobstore.BlobManagerInterface = new(FakeBlobManagerInterface)
