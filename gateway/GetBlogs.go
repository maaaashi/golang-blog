package gateway

import "go-blog/domain"

type BlogJson struct {
	Id      int
	Title   string
	Content string
}

type GetBlogsInputPort interface {
	GetBlogs() ([]BlogJson, error)
}

type GetBlogsGateway struct {
	BlogsDriver GetBlogsInputPort
}

func NewGetBlogsGateway(blogsDriver GetBlogsInputPort) *GetBlogsGateway {
	return &GetBlogsGateway{
		BlogsDriver: blogsDriver,
	}
}

func (g *GetBlogsGateway) Execute() ([]domain.Blog, error) {
	var blogs []domain.Blog
	blogsJson, _ := g.BlogsDriver.GetBlogs()

	for _, blogJson := range blogsJson {
		blogs = append(blogs, domain.Blog{Id: domain.BlogId(blogJson.Id), Title: domain.BlogTItle(blogJson.Title), Content: domain.BlogContent(blogJson.Content)})
	}
	return blogs, nil
}
