package usecase

import (
	"go-blog/domain"
)

type GetBlogsInputPort interface {
	Execute() ([]domain.Blog, error)
}

type GetBlogsUsecase struct {
	GetBlogsInputPort GetBlogsInputPort
}

func NewGetBlogsUsecase(inputPort GetBlogsInputPort) GetBlogsUsecase {
	return GetBlogsUsecase{
		GetBlogsInputPort: inputPort,
	}
}

func (g *GetBlogsUsecase) Execute() ([]domain.Blog, error) {
	return g.GetBlogsInputPort.Execute()
}
