package mocks

import (
	entity "github.com/sawalreverr/recything/internal/feature/article/entity"
	pagination "github.com/sawalreverr/recything/pagination"
	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

type ArticleRepositoryInterface struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: articleInput, image
func (_m *ArticleRepositoryInterface) CreateArticle(articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	ret := _m.Called(articleInput, image)

	if len(ret) == 0 {
		panic("no return value specified for CreateArticle")
	}

	var r0 entity.ArticleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.ArticleCore, *multipart.FileHeader) (entity.ArticleCore, error)); ok {
		return rf(articleInput, image)
	}
	if rf, ok := ret.Get(0).(func(entity.ArticleCore, *multipart.FileHeader) entity.ArticleCore); ok {
		r0 = rf(articleInput, image)
	} else {
		r0 = ret.Get(0).(entity.ArticleCore)
	}

	if rf, ok := ret.Get(1).(func(entity.ArticleCore, *multipart.FileHeader) error); ok {
		r1 = rf(articleInput, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteArticle provides a mock function with given fields: id
func (_m *ArticleRepositoryInterface) DeleteArticle(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllArticle provides a mock function with given fields: page, limit, search, filter
func (_m *ArticleRepositoryInterface) GetAllArticle(page int, limit int, search string, filter string) ([]entity.ArticleCore, pagination.Pageinfo, int, error) {
	ret := _m.Called(page, limit, search, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAllArticle")
	}

	var r0 []entity.ArticleCore
	var r1 pagination.Pageinfo
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]entity.ArticleCore, pagination.Pageinfo, int, error)); ok {
		return rf(page, limit, search, filter)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []entity.ArticleCore); ok {
		r0 = rf(page, limit, search, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ArticleCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) pagination.Pageinfo); ok {
		r1 = rf(page, limit, search, filter)
	} else {
		r1 = ret.Get(1).(pagination.Pageinfo)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) int); ok {
		r2 = rf(page, limit, search, filter)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(int, int, string, string) error); ok {
		r3 = rf(page, limit, search, filter)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetPopularArticle provides a mock function with given fields: search
func (_m *ArticleRepositoryInterface) GetPopularArticle(search string) ([]entity.ArticleCore, error) {
	ret := _m.Called(search)

	if len(ret) == 0 {
		panic("no return value specified for GetPopularArticle")
	}

	var r0 []entity.ArticleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entity.ArticleCore, error)); ok {
		return rf(search)
	}
	if rf, ok := ret.Get(0).(func(string) []entity.ArticleCore); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ArticleCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpecificArticle provides a mock function with given fields: idArticle
func (_m *ArticleRepositoryInterface) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	ret := _m.Called(idArticle)

	if len(ret) == 0 {
		panic("no return value specified for GetSpecificArticle")
	}

	var r0 entity.ArticleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.ArticleCore, error)); ok {
		return rf(idArticle)
	}
	if rf, ok := ret.Get(0).(func(string) entity.ArticleCore); ok {
		r0 = rf(idArticle)
	} else {
		r0 = ret.Get(0).(entity.ArticleCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idArticle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArticle provides a mock function with given fields: idArticle, articleInput, image
func (_m *ArticleRepositoryInterface) UpdateArticle(idArticle string, articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	ret := _m.Called(idArticle, articleInput, image)

	if len(ret) == 0 {
		panic("no return value specified for UpdateArticle")
	}

	var r0 entity.ArticleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, entity.ArticleCore, *multipart.FileHeader) (entity.ArticleCore, error)); ok {
		return rf(idArticle, articleInput, image)
	}
	if rf, ok := ret.Get(0).(func(string, entity.ArticleCore, *multipart.FileHeader) entity.ArticleCore); ok {
		r0 = rf(idArticle, articleInput, image)
	} else {
		r0 = ret.Get(0).(entity.ArticleCore)
	}

	if rf, ok := ret.Get(1).(func(string, entity.ArticleCore, *multipart.FileHeader) error); ok {
		r1 = rf(idArticle, articleInput, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewArticleRepositoryInterface creates a new instance of ArticleRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleRepositoryInterface {
	mock := &ArticleRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
