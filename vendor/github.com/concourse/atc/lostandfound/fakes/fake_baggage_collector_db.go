// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/lostandfound"
)

type FakeBaggageCollectorDB struct {
	ReapVolumeStub        func(string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		arg1 string
	}
	reapVolumeReturns struct {
		result1 error
	}
	GetAllPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPipelinesMutex       sync.RWMutex
	getAllPipelinesArgsForCall []struct{}
	getAllPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetVolumesStub        func() ([]db.SavedVolume, error)
	getVolumesMutex       sync.RWMutex
	getVolumesArgsForCall []struct{}
	getVolumesReturns     struct {
		result1 []db.SavedVolume
		result2 error
	}
	GetImageResourceCacheIdentifiersByBuildIDStub        func(buildID int) ([]db.ResourceCacheIdentifier, error)
	getImageResourceCacheIdentifiersByBuildIDMutex       sync.RWMutex
	getImageResourceCacheIdentifiersByBuildIDArgsForCall []struct {
		buildID int
	}
	getImageResourceCacheIdentifiersByBuildIDReturns struct {
		result1 []db.ResourceCacheIdentifier
		result2 error
	}
	GetVolumesForOneOffBuildImageResourcesStub        func() ([]db.SavedVolume, error)
	getVolumesForOneOffBuildImageResourcesMutex       sync.RWMutex
	getVolumesForOneOffBuildImageResourcesArgsForCall []struct{}
	getVolumesForOneOffBuildImageResourcesReturns     struct {
		result1 []db.SavedVolume
		result2 error
	}
}

func (fake *FakeBaggageCollectorDB) ReapVolume(arg1 string) error {
	fake.reapVolumeMutex.Lock()
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(arg1)
	} else {
		return fake.reapVolumeReturns.result1
	}
}

func (fake *FakeBaggageCollectorDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeBaggageCollectorDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].arg1
}

func (fake *FakeBaggageCollectorDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBaggageCollectorDB) GetAllPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPipelinesMutex.Lock()
	fake.getAllPipelinesArgsForCall = append(fake.getAllPipelinesArgsForCall, struct{}{})
	fake.getAllPipelinesMutex.Unlock()
	if fake.GetAllPipelinesStub != nil {
		return fake.GetAllPipelinesStub()
	} else {
		return fake.getAllPipelinesReturns.result1, fake.getAllPipelinesReturns.result2
	}
}

func (fake *FakeBaggageCollectorDB) GetAllPipelinesCallCount() int {
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return len(fake.getAllPipelinesArgsForCall)
}

func (fake *FakeBaggageCollectorDB) GetAllPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	fake.getAllPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeBaggageCollectorDB) GetVolumes() ([]db.SavedVolume, error) {
	fake.getVolumesMutex.Lock()
	fake.getVolumesArgsForCall = append(fake.getVolumesArgsForCall, struct{}{})
	fake.getVolumesMutex.Unlock()
	if fake.GetVolumesStub != nil {
		return fake.GetVolumesStub()
	} else {
		return fake.getVolumesReturns.result1, fake.getVolumesReturns.result2
	}
}

func (fake *FakeBaggageCollectorDB) GetVolumesCallCount() int {
	fake.getVolumesMutex.RLock()
	defer fake.getVolumesMutex.RUnlock()
	return len(fake.getVolumesArgsForCall)
}

func (fake *FakeBaggageCollectorDB) GetVolumesReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesStub = nil
	fake.getVolumesReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeBaggageCollectorDB) GetImageResourceCacheIdentifiersByBuildID(buildID int) ([]db.ResourceCacheIdentifier, error) {
	fake.getImageResourceCacheIdentifiersByBuildIDMutex.Lock()
	fake.getImageResourceCacheIdentifiersByBuildIDArgsForCall = append(fake.getImageResourceCacheIdentifiersByBuildIDArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getImageResourceCacheIdentifiersByBuildIDMutex.Unlock()
	if fake.GetImageResourceCacheIdentifiersByBuildIDStub != nil {
		return fake.GetImageResourceCacheIdentifiersByBuildIDStub(buildID)
	} else {
		return fake.getImageResourceCacheIdentifiersByBuildIDReturns.result1, fake.getImageResourceCacheIdentifiersByBuildIDReturns.result2
	}
}

func (fake *FakeBaggageCollectorDB) GetImageResourceCacheIdentifiersByBuildIDCallCount() int {
	fake.getImageResourceCacheIdentifiersByBuildIDMutex.RLock()
	defer fake.getImageResourceCacheIdentifiersByBuildIDMutex.RUnlock()
	return len(fake.getImageResourceCacheIdentifiersByBuildIDArgsForCall)
}

func (fake *FakeBaggageCollectorDB) GetImageResourceCacheIdentifiersByBuildIDArgsForCall(i int) int {
	fake.getImageResourceCacheIdentifiersByBuildIDMutex.RLock()
	defer fake.getImageResourceCacheIdentifiersByBuildIDMutex.RUnlock()
	return fake.getImageResourceCacheIdentifiersByBuildIDArgsForCall[i].buildID
}

func (fake *FakeBaggageCollectorDB) GetImageResourceCacheIdentifiersByBuildIDReturns(result1 []db.ResourceCacheIdentifier, result2 error) {
	fake.GetImageResourceCacheIdentifiersByBuildIDStub = nil
	fake.getImageResourceCacheIdentifiersByBuildIDReturns = struct {
		result1 []db.ResourceCacheIdentifier
		result2 error
	}{result1, result2}
}

func (fake *FakeBaggageCollectorDB) GetVolumesForOneOffBuildImageResources() ([]db.SavedVolume, error) {
	fake.getVolumesForOneOffBuildImageResourcesMutex.Lock()
	fake.getVolumesForOneOffBuildImageResourcesArgsForCall = append(fake.getVolumesForOneOffBuildImageResourcesArgsForCall, struct{}{})
	fake.getVolumesForOneOffBuildImageResourcesMutex.Unlock()
	if fake.GetVolumesForOneOffBuildImageResourcesStub != nil {
		return fake.GetVolumesForOneOffBuildImageResourcesStub()
	} else {
		return fake.getVolumesForOneOffBuildImageResourcesReturns.result1, fake.getVolumesForOneOffBuildImageResourcesReturns.result2
	}
}

func (fake *FakeBaggageCollectorDB) GetVolumesForOneOffBuildImageResourcesCallCount() int {
	fake.getVolumesForOneOffBuildImageResourcesMutex.RLock()
	defer fake.getVolumesForOneOffBuildImageResourcesMutex.RUnlock()
	return len(fake.getVolumesForOneOffBuildImageResourcesArgsForCall)
}

func (fake *FakeBaggageCollectorDB) GetVolumesForOneOffBuildImageResourcesReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesForOneOffBuildImageResourcesStub = nil
	fake.getVolumesForOneOffBuildImageResourcesReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

var _ lostandfound.BaggageCollectorDB = new(FakeBaggageCollectorDB)