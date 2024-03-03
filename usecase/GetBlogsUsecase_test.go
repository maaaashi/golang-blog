package usecase_test

import (
	"go-blog/domain"
	"go-blog/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockGetBlogsInputPort struct {
	mock.Mock
}

func (m *MockGetBlogsInputPort) Execute() ([]*domain.Blog, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Blog), args.Error(1)
}

func TestUsecaseExecute(t *testing.T) {
	mockInputPort := new(MockGetBlogsInputPort)
	mockInputPort.On("Execute").Return([]*domain.Blog{
		{Id: domain.BlogId(1), Title: domain.BlogTItle("title"), Content: domain.BlogContent("content")},
	}, nil)
	usecase := usecase.NewGetBlogsUsecase(mockInputPort)

	actual, _ := usecase.Execute()
	expected := []*domain.Blog{{Id: domain.BlogId(1), Title: domain.BlogTItle("title"), Content: domain.BlogContent("content")}}

	if actual[0].Id != expected[0].Id {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	mockInputPort.AssertCalled(t, "Execute")
}
