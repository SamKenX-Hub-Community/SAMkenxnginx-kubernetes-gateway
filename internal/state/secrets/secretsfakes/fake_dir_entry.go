// Code generated by counterfeiter. DO NOT EDIT.
package secretsfakes

import (
	"io/fs"
	"sync"
)

type FakeDirEntry struct {
	InfoStub        func() (fs.FileInfo, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
	}
	infoReturns struct {
		result1 fs.FileInfo
		result2 error
	}
	infoReturnsOnCall map[int]struct {
		result1 fs.FileInfo
		result2 error
	}
	IsDirStub        func() bool
	isDirMutex       sync.RWMutex
	isDirArgsForCall []struct {
	}
	isDirReturns struct {
		result1 bool
	}
	isDirReturnsOnCall map[int]struct {
		result1 bool
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() fs.FileMode
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 fs.FileMode
	}
	typeReturnsOnCall map[int]struct {
		result1 fs.FileMode
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirEntry) Info() (fs.FileInfo, error) {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
	}{})
	stub := fake.InfoStub
	fakeReturns := fake.infoReturns
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDirEntry) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeDirEntry) InfoCalls(stub func() (fs.FileInfo, error)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *FakeDirEntry) InfoReturns(result1 fs.FileInfo, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDirEntry) InfoReturnsOnCall(i int, result1 fs.FileInfo, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 fs.FileInfo
			result2 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDirEntry) IsDir() bool {
	fake.isDirMutex.Lock()
	ret, specificReturn := fake.isDirReturnsOnCall[len(fake.isDirArgsForCall)]
	fake.isDirArgsForCall = append(fake.isDirArgsForCall, struct {
	}{})
	stub := fake.IsDirStub
	fakeReturns := fake.isDirReturns
	fake.recordInvocation("IsDir", []interface{}{})
	fake.isDirMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDirEntry) IsDirCallCount() int {
	fake.isDirMutex.RLock()
	defer fake.isDirMutex.RUnlock()
	return len(fake.isDirArgsForCall)
}

func (fake *FakeDirEntry) IsDirCalls(stub func() bool) {
	fake.isDirMutex.Lock()
	defer fake.isDirMutex.Unlock()
	fake.IsDirStub = stub
}

func (fake *FakeDirEntry) IsDirReturns(result1 bool) {
	fake.isDirMutex.Lock()
	defer fake.isDirMutex.Unlock()
	fake.IsDirStub = nil
	fake.isDirReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDirEntry) IsDirReturnsOnCall(i int, result1 bool) {
	fake.isDirMutex.Lock()
	defer fake.isDirMutex.Unlock()
	fake.IsDirStub = nil
	if fake.isDirReturnsOnCall == nil {
		fake.isDirReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isDirReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDirEntry) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDirEntry) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeDirEntry) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeDirEntry) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeDirEntry) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeDirEntry) Type() fs.FileMode {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	stub := fake.TypeStub
	fakeReturns := fake.typeReturns
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDirEntry) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeDirEntry) TypeCalls(stub func() fs.FileMode) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeDirEntry) TypeReturns(result1 fs.FileMode) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 fs.FileMode
	}{result1}
}

func (fake *FakeDirEntry) TypeReturnsOnCall(i int, result1 fs.FileMode) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 fs.FileMode
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 fs.FileMode
	}{result1}
}

func (fake *FakeDirEntry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.isDirMutex.RLock()
	defer fake.isDirMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDirEntry) recordInvocation(key string, args []interface{}) {
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

var _ fs.DirEntry = new(FakeDirEntry)
