// This file was generated by counterfeiter
package filegroupfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/go-pivnet/cmd/pivnet/commands/filegroup"
)

type FakePivnetClient struct {
	FileGroupsStub        func(productSlug string) ([]go_pivnet.FileGroup, error)
	fileGroupsMutex       sync.RWMutex
	fileGroupsArgsForCall []struct {
		productSlug string
	}
	fileGroupsReturns struct {
		result1 []go_pivnet.FileGroup
		result2 error
	}
	FileGroupsForReleaseStub        func(productSlug string, releaseID int) ([]go_pivnet.FileGroup, error)
	fileGroupsForReleaseMutex       sync.RWMutex
	fileGroupsForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	fileGroupsForReleaseReturns struct {
		result1 []go_pivnet.FileGroup
		result2 error
	}
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	FileGroupStub        func(productSlug string, fileGroupID int) (go_pivnet.FileGroup, error)
	fileGroupMutex       sync.RWMutex
	fileGroupArgsForCall []struct {
		productSlug string
		fileGroupID int
	}
	fileGroupReturns struct {
		result1 go_pivnet.FileGroup
		result2 error
	}
	CreateFileGroupStub        func(productSlug string, name string) (go_pivnet.FileGroup, error)
	createFileGroupMutex       sync.RWMutex
	createFileGroupArgsForCall []struct {
		productSlug string
		name        string
	}
	createFileGroupReturns struct {
		result1 go_pivnet.FileGroup
		result2 error
	}
	UpdateFileGroupStub        func(productSlug string, fileGroup go_pivnet.FileGroup) (go_pivnet.FileGroup, error)
	updateFileGroupMutex       sync.RWMutex
	updateFileGroupArgsForCall []struct {
		productSlug string
		fileGroup   go_pivnet.FileGroup
	}
	updateFileGroupReturns struct {
		result1 go_pivnet.FileGroup
		result2 error
	}
	DeleteFileGroupStub        func(productSlug string, fileGroupID int) (go_pivnet.FileGroup, error)
	deleteFileGroupMutex       sync.RWMutex
	deleteFileGroupArgsForCall []struct {
		productSlug string
		fileGroupID int
	}
	deleteFileGroupReturns struct {
		result1 go_pivnet.FileGroup
		result2 error
	}
	AddFileGroupToReleaseStub        func(productSlug string, fileGroupID int, releaseID int) error
	addFileGroupToReleaseMutex       sync.RWMutex
	addFileGroupToReleaseArgsForCall []struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}
	addFileGroupToReleaseReturns struct {
		result1 error
	}
	RemoveFileGroupFromReleaseStub        func(productSlug string, fileGroupID int, releaseID int) error
	removeFileGroupFromReleaseMutex       sync.RWMutex
	removeFileGroupFromReleaseArgsForCall []struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}
	removeFileGroupFromReleaseReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) FileGroups(productSlug string) ([]go_pivnet.FileGroup, error) {
	fake.fileGroupsMutex.Lock()
	fake.fileGroupsArgsForCall = append(fake.fileGroupsArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("FileGroups", []interface{}{productSlug})
	fake.fileGroupsMutex.Unlock()
	if fake.FileGroupsStub != nil {
		return fake.FileGroupsStub(productSlug)
	} else {
		return fake.fileGroupsReturns.result1, fake.fileGroupsReturns.result2
	}
}

func (fake *FakePivnetClient) FileGroupsCallCount() int {
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	return len(fake.fileGroupsArgsForCall)
}

func (fake *FakePivnetClient) FileGroupsArgsForCall(i int) string {
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	return fake.fileGroupsArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) FileGroupsReturns(result1 []go_pivnet.FileGroup, result2 error) {
	fake.FileGroupsStub = nil
	fake.fileGroupsReturns = struct {
		result1 []go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupsForRelease(productSlug string, releaseID int) ([]go_pivnet.FileGroup, error) {
	fake.fileGroupsForReleaseMutex.Lock()
	fake.fileGroupsForReleaseArgsForCall = append(fake.fileGroupsForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("FileGroupsForRelease", []interface{}{productSlug, releaseID})
	fake.fileGroupsForReleaseMutex.Unlock()
	if fake.FileGroupsForReleaseStub != nil {
		return fake.FileGroupsForReleaseStub(productSlug, releaseID)
	} else {
		return fake.fileGroupsForReleaseReturns.result1, fake.fileGroupsForReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) FileGroupsForReleaseCallCount() int {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return len(fake.fileGroupsForReleaseArgsForCall)
}

func (fake *FakePivnetClient) FileGroupsForReleaseArgsForCall(i int) (string, int) {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return fake.fileGroupsForReleaseArgsForCall[i].productSlug, fake.fileGroupsForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) FileGroupsForReleaseReturns(result1 []go_pivnet.FileGroup, result2 error) {
	fake.FileGroupsForReleaseStub = nil
	fake.fileGroupsForReleaseReturns = struct {
		result1 []go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 go_pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroup(productSlug string, fileGroupID int) (go_pivnet.FileGroup, error) {
	fake.fileGroupMutex.Lock()
	fake.fileGroupArgsForCall = append(fake.fileGroupArgsForCall, struct {
		productSlug string
		fileGroupID int
	}{productSlug, fileGroupID})
	fake.recordInvocation("FileGroup", []interface{}{productSlug, fileGroupID})
	fake.fileGroupMutex.Unlock()
	if fake.FileGroupStub != nil {
		return fake.FileGroupStub(productSlug, fileGroupID)
	} else {
		return fake.fileGroupReturns.result1, fake.fileGroupReturns.result2
	}
}

func (fake *FakePivnetClient) FileGroupCallCount() int {
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	return len(fake.fileGroupArgsForCall)
}

func (fake *FakePivnetClient) FileGroupArgsForCall(i int) (string, int) {
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	return fake.fileGroupArgsForCall[i].productSlug, fake.fileGroupArgsForCall[i].fileGroupID
}

func (fake *FakePivnetClient) FileGroupReturns(result1 go_pivnet.FileGroup, result2 error) {
	fake.FileGroupStub = nil
	fake.fileGroupReturns = struct {
		result1 go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateFileGroup(productSlug string, name string) (go_pivnet.FileGroup, error) {
	fake.createFileGroupMutex.Lock()
	fake.createFileGroupArgsForCall = append(fake.createFileGroupArgsForCall, struct {
		productSlug string
		name        string
	}{productSlug, name})
	fake.recordInvocation("CreateFileGroup", []interface{}{productSlug, name})
	fake.createFileGroupMutex.Unlock()
	if fake.CreateFileGroupStub != nil {
		return fake.CreateFileGroupStub(productSlug, name)
	} else {
		return fake.createFileGroupReturns.result1, fake.createFileGroupReturns.result2
	}
}

func (fake *FakePivnetClient) CreateFileGroupCallCount() int {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return len(fake.createFileGroupArgsForCall)
}

func (fake *FakePivnetClient) CreateFileGroupArgsForCall(i int) (string, string) {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return fake.createFileGroupArgsForCall[i].productSlug, fake.createFileGroupArgsForCall[i].name
}

func (fake *FakePivnetClient) CreateFileGroupReturns(result1 go_pivnet.FileGroup, result2 error) {
	fake.CreateFileGroupStub = nil
	fake.createFileGroupReturns = struct {
		result1 go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateFileGroup(productSlug string, fileGroup go_pivnet.FileGroup) (go_pivnet.FileGroup, error) {
	fake.updateFileGroupMutex.Lock()
	fake.updateFileGroupArgsForCall = append(fake.updateFileGroupArgsForCall, struct {
		productSlug string
		fileGroup   go_pivnet.FileGroup
	}{productSlug, fileGroup})
	fake.recordInvocation("UpdateFileGroup", []interface{}{productSlug, fileGroup})
	fake.updateFileGroupMutex.Unlock()
	if fake.UpdateFileGroupStub != nil {
		return fake.UpdateFileGroupStub(productSlug, fileGroup)
	} else {
		return fake.updateFileGroupReturns.result1, fake.updateFileGroupReturns.result2
	}
}

func (fake *FakePivnetClient) UpdateFileGroupCallCount() int {
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	return len(fake.updateFileGroupArgsForCall)
}

func (fake *FakePivnetClient) UpdateFileGroupArgsForCall(i int) (string, go_pivnet.FileGroup) {
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	return fake.updateFileGroupArgsForCall[i].productSlug, fake.updateFileGroupArgsForCall[i].fileGroup
}

func (fake *FakePivnetClient) UpdateFileGroupReturns(result1 go_pivnet.FileGroup, result2 error) {
	fake.UpdateFileGroupStub = nil
	fake.updateFileGroupReturns = struct {
		result1 go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteFileGroup(productSlug string, fileGroupID int) (go_pivnet.FileGroup, error) {
	fake.deleteFileGroupMutex.Lock()
	fake.deleteFileGroupArgsForCall = append(fake.deleteFileGroupArgsForCall, struct {
		productSlug string
		fileGroupID int
	}{productSlug, fileGroupID})
	fake.recordInvocation("DeleteFileGroup", []interface{}{productSlug, fileGroupID})
	fake.deleteFileGroupMutex.Unlock()
	if fake.DeleteFileGroupStub != nil {
		return fake.DeleteFileGroupStub(productSlug, fileGroupID)
	} else {
		return fake.deleteFileGroupReturns.result1, fake.deleteFileGroupReturns.result2
	}
}

func (fake *FakePivnetClient) DeleteFileGroupCallCount() int {
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	return len(fake.deleteFileGroupArgsForCall)
}

func (fake *FakePivnetClient) DeleteFileGroupArgsForCall(i int) (string, int) {
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	return fake.deleteFileGroupArgsForCall[i].productSlug, fake.deleteFileGroupArgsForCall[i].fileGroupID
}

func (fake *FakePivnetClient) DeleteFileGroupReturns(result1 go_pivnet.FileGroup, result2 error) {
	fake.DeleteFileGroupStub = nil
	fake.deleteFileGroupReturns = struct {
		result1 go_pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddFileGroupToRelease(productSlug string, fileGroupID int, releaseID int) error {
	fake.addFileGroupToReleaseMutex.Lock()
	fake.addFileGroupToReleaseArgsForCall = append(fake.addFileGroupToReleaseArgsForCall, struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}{productSlug, fileGroupID, releaseID})
	fake.recordInvocation("AddFileGroupToRelease", []interface{}{productSlug, fileGroupID, releaseID})
	fake.addFileGroupToReleaseMutex.Unlock()
	if fake.AddFileGroupToReleaseStub != nil {
		return fake.AddFileGroupToReleaseStub(productSlug, fileGroupID, releaseID)
	} else {
		return fake.addFileGroupToReleaseReturns.result1
	}
}

func (fake *FakePivnetClient) AddFileGroupToReleaseCallCount() int {
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	return len(fake.addFileGroupToReleaseArgsForCall)
}

func (fake *FakePivnetClient) AddFileGroupToReleaseArgsForCall(i int) (string, int, int) {
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	return fake.addFileGroupToReleaseArgsForCall[i].productSlug, fake.addFileGroupToReleaseArgsForCall[i].fileGroupID, fake.addFileGroupToReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AddFileGroupToReleaseReturns(result1 error) {
	fake.AddFileGroupToReleaseStub = nil
	fake.addFileGroupToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveFileGroupFromRelease(productSlug string, fileGroupID int, releaseID int) error {
	fake.removeFileGroupFromReleaseMutex.Lock()
	fake.removeFileGroupFromReleaseArgsForCall = append(fake.removeFileGroupFromReleaseArgsForCall, struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}{productSlug, fileGroupID, releaseID})
	fake.recordInvocation("RemoveFileGroupFromRelease", []interface{}{productSlug, fileGroupID, releaseID})
	fake.removeFileGroupFromReleaseMutex.Unlock()
	if fake.RemoveFileGroupFromReleaseStub != nil {
		return fake.RemoveFileGroupFromReleaseStub(productSlug, fileGroupID, releaseID)
	} else {
		return fake.removeFileGroupFromReleaseReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseCallCount() int {
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	return len(fake.removeFileGroupFromReleaseArgsForCall)
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseArgsForCall(i int) (string, int, int) {
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	return fake.removeFileGroupFromReleaseArgsForCall[i].productSlug, fake.removeFileGroupFromReleaseArgsForCall[i].fileGroupID, fake.removeFileGroupFromReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseReturns(result1 error) {
	fake.RemoveFileGroupFromReleaseStub = nil
	fake.removeFileGroupFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ filegroup.PivnetClient = new(FakePivnetClient)
