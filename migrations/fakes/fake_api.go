// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cfmobile/gopivnet/api"
	"github.com/cfmobile/gopivnet/resource"
)

type FakeApi struct {
	GetLatestProductFileStub        func(productName string, fileType string) (*resource.ProductFile, error)
	getLatestProductFileMutex       sync.RWMutex
	getLatestProductFileArgsForCall []struct {
		productName string
		fileType    string
	}
	getLatestProductFileReturns struct {
		result1 *resource.ProductFile
		result2 error
	}
	GetProductFileForVersionStub        func(productName, version string, fileType string) (*resource.ProductFile, error)
	getProductFileForVersionMutex       sync.RWMutex
	getProductFileForVersionArgsForCall []struct {
		productName string
		version     string
		fileType    string
	}
	getProductFileForVersionReturns struct {
		result1 *resource.ProductFile
		result2 error
	}
	GetVersionsForProductStub        func(productName string) ([]string, error)
	getVersionsForProductMutex       sync.RWMutex
	getVersionsForProductArgsForCall []struct {
		productName string
	}
	getVersionsForProductReturns struct {
		result1 []string
		result2 error
	}
	DownloadStub        func(productFile *resource.ProductFile, fileName string) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		productFile *resource.ProductFile
		fileName    string
	}
	downloadReturns struct {
		result1 error
	}
}

func (fake *FakeApi) GetLatestProductFile(productName string, fileType string) (*resource.ProductFile, error) {
	fake.getLatestProductFileMutex.Lock()
	fake.getLatestProductFileArgsForCall = append(fake.getLatestProductFileArgsForCall, struct {
		productName string
		fileType    string
	}{productName, fileType})
	fake.getLatestProductFileMutex.Unlock()
	if fake.GetLatestProductFileStub != nil {
		return fake.GetLatestProductFileStub(productName, fileType)
	} else {
		return fake.getLatestProductFileReturns.result1, fake.getLatestProductFileReturns.result2
	}
}

func (fake *FakeApi) GetLatestProductFileCallCount() int {
	fake.getLatestProductFileMutex.RLock()
	defer fake.getLatestProductFileMutex.RUnlock()
	return len(fake.getLatestProductFileArgsForCall)
}

func (fake *FakeApi) GetLatestProductFileArgsForCall(i int) (string, string) {
	fake.getLatestProductFileMutex.RLock()
	defer fake.getLatestProductFileMutex.RUnlock()
	return fake.getLatestProductFileArgsForCall[i].productName, fake.getLatestProductFileArgsForCall[i].fileType
}

func (fake *FakeApi) GetLatestProductFileReturns(result1 *resource.ProductFile, result2 error) {
	fake.GetLatestProductFileStub = nil
	fake.getLatestProductFileReturns = struct {
		result1 *resource.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeApi) GetProductFileForVersion(productName string, version string, fileType string) (*resource.ProductFile, error) {
	fake.getProductFileForVersionMutex.Lock()
	fake.getProductFileForVersionArgsForCall = append(fake.getProductFileForVersionArgsForCall, struct {
		productName string
		version     string
		fileType    string
	}{productName, version, fileType})
	fake.getProductFileForVersionMutex.Unlock()
	if fake.GetProductFileForVersionStub != nil {
		return fake.GetProductFileForVersionStub(productName, version, fileType)
	} else {
		return fake.getProductFileForVersionReturns.result1, fake.getProductFileForVersionReturns.result2
	}
}

func (fake *FakeApi) GetProductFileForVersionCallCount() int {
	fake.getProductFileForVersionMutex.RLock()
	defer fake.getProductFileForVersionMutex.RUnlock()
	return len(fake.getProductFileForVersionArgsForCall)
}

func (fake *FakeApi) GetProductFileForVersionArgsForCall(i int) (string, string, string) {
	fake.getProductFileForVersionMutex.RLock()
	defer fake.getProductFileForVersionMutex.RUnlock()
	return fake.getProductFileForVersionArgsForCall[i].productName, fake.getProductFileForVersionArgsForCall[i].version, fake.getProductFileForVersionArgsForCall[i].fileType
}

func (fake *FakeApi) GetProductFileForVersionReturns(result1 *resource.ProductFile, result2 error) {
	fake.GetProductFileForVersionStub = nil
	fake.getProductFileForVersionReturns = struct {
		result1 *resource.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeApi) GetVersionsForProduct(productName string) ([]string, error) {
	fake.getVersionsForProductMutex.Lock()
	fake.getVersionsForProductArgsForCall = append(fake.getVersionsForProductArgsForCall, struct {
		productName string
	}{productName})
	fake.getVersionsForProductMutex.Unlock()
	if fake.GetVersionsForProductStub != nil {
		return fake.GetVersionsForProductStub(productName)
	} else {
		return fake.getVersionsForProductReturns.result1, fake.getVersionsForProductReturns.result2
	}
}

func (fake *FakeApi) GetVersionsForProductCallCount() int {
	fake.getVersionsForProductMutex.RLock()
	defer fake.getVersionsForProductMutex.RUnlock()
	return len(fake.getVersionsForProductArgsForCall)
}

func (fake *FakeApi) GetVersionsForProductArgsForCall(i int) string {
	fake.getVersionsForProductMutex.RLock()
	defer fake.getVersionsForProductMutex.RUnlock()
	return fake.getVersionsForProductArgsForCall[i].productName
}

func (fake *FakeApi) GetVersionsForProductReturns(result1 []string, result2 error) {
	fake.GetVersionsForProductStub = nil
	fake.getVersionsForProductReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApi) Download(productFile *resource.ProductFile, fileName string) error {
	fake.downloadMutex.Lock()
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		productFile *resource.ProductFile
		fileName    string
	}{productFile, fileName})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(productFile, fileName)
	} else {
		return fake.downloadReturns.result1
	}
}

func (fake *FakeApi) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeApi) DownloadArgsForCall(i int) (*resource.ProductFile, string) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.downloadArgsForCall[i].productFile, fake.downloadArgsForCall[i].fileName
}

func (fake *FakeApi) DownloadReturns(result1 error) {
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

var _ api.Api = new(FakeApi)
