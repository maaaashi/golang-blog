package gateway_test

import (
	"go-blog/domain"
	"go-blog/gateway"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockBlogDriver struct {
	mock.Mock
}

func (m *MockBlogDriver) GetBlogs() ([]*gateway.BlogJson, error) {
	args := m.Called()
	return args.Get(0).([]*gateway.BlogJson), args.Error(1)
}

func TestGatewayExecute(t *testing.T) {
	mockDriver := new(MockBlogDriver)
	mockDriver.On("GetBlogs").Return([]*gateway.BlogJson{
		{Id: 1, Title: "title", Content: "content"},
	}, nil)

	gateway := gateway.NewGetBlogsGateway(mockDriver)
	actual, _ := gateway.Execute()
	expected := []*domain.Blog{
		{Id: domain.BlogId(1), Title: domain.BlogTItle("title"), Content: domain.BlogContent("content")},
	}

	if actual[0].Id != expected[0].Id {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	mockDriver.AssertCalled(t, "GetBlogs")
}
