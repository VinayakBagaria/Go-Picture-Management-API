package tests

import (
	"mime/multipart"
	"reflect"
	"strings"
	"testing"

	"github.com/VinayakBagaria/go-cat-pictures/service"
)

func newFile(fileName string) *multipart.FileHeader {
	return &multipart.FileHeader{
		Filename: fileName,
		Size:     1000,
	}
}

func TestApiHandlers(t *testing.T) {
	fakeRepo := &FakeRepository{}
	fakeStorage := NewFakeStorage()
	svc := service.NewPicturesService(fakeRepo, fakeStorage)

	t.Run("create entry", func(t *testing.T) {
		file := newFile(NewUniqueString())
		createResponse, errorState := svc.Create(file)
		if errorState != nil {
			assertNoError(t, errorState.Error)
		}

		assertData(t, true, strings.HasSuffix(createResponse.Name, file.Filename))
		assertData(t, int(createResponse.Id), 1)

		fileResponse, err := svc.Get(int(createResponse.Id))
		assertNoError(t, err)
		assertData(t, fileResponse.Name, createResponse.Name)
	})

	t.Run("update entry", func(t *testing.T) {
		file := newFile(NewUniqueString())
		updateResponse, errorState := svc.Update(1, file)
		if errorState != nil {
			assertNoError(t, errorState.Error)
		}

		assertData(t, true, strings.HasSuffix(updateResponse.Name, file.Filename))

		fileResponse, err := svc.Get(int(updateResponse.Id))
		assertNoError(t, err)
		assertData(t, fileResponse.Name, updateResponse.Name)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertData(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
