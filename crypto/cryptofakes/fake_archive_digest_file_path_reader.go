// This file was generated by counterfeiter
package cryptofakes

import (
	"os"
	"sync"

	"github.com/bluebosh/bosh-utils/crypto"
	"github.com/bluebosh/bosh-utils/system"
)

type FakeArchiveDigestFilePathReader struct {
	OpenFileStub        func(path string, flag int, perm os.FileMode) (system.File, error)
	openFileMutex       sync.RWMutex
	openFileArgsForCall []struct {
		path string
		flag int
		perm os.FileMode
	}
	openFileReturns struct {
		result1 system.File
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArchiveDigestFilePathReader) OpenFile(path string, flag int, perm os.FileMode) (system.File, error) {
	fake.openFileMutex.Lock()
	fake.openFileArgsForCall = append(fake.openFileArgsForCall, struct {
		path string
		flag int
		perm os.FileMode
	}{path, flag, perm})
	fake.recordInvocation("OpenFile", []interface{}{path, flag, perm})
	fake.openFileMutex.Unlock()
	if fake.OpenFileStub != nil {
		return fake.OpenFileStub(path, flag, perm)
	} else {
		return fake.openFileReturns.result1, fake.openFileReturns.result2
	}
}

func (fake *FakeArchiveDigestFilePathReader) OpenFileCallCount() int {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return len(fake.openFileArgsForCall)
}

func (fake *FakeArchiveDigestFilePathReader) OpenFileArgsForCall(i int) (string, int, os.FileMode) {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.openFileArgsForCall[i].path, fake.openFileArgsForCall[i].flag, fake.openFileArgsForCall[i].perm
}

func (fake *FakeArchiveDigestFilePathReader) OpenFileReturns(result1 system.File, result2 error) {
	fake.OpenFileStub = nil
	fake.openFileReturns = struct {
		result1 system.File
		result2 error
	}{result1, result2}
}

func (fake *FakeArchiveDigestFilePathReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeArchiveDigestFilePathReader) recordInvocation(key string, args []interface{}) {
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

var _ crypto.ArchiveDigestFilePathReader = new(FakeArchiveDigestFilePathReader)
