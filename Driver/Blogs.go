package driver

import "go-blog/gateway"

type BlogsDriver struct {
	// Urlとかの情報を持つ
}

func NewBlogsDriver() *BlogsDriver {
	return &BlogsDriver{}
}

func (d *BlogsDriver) GetBlogs() ([]gateway.BlogJson, error) {
	return []gateway.BlogJson{
		{
			Id:      1,
			Title:   "Title 1",
			Content: "Content 1",
		},
		{
			Id:      2,
			Title:   "Title 2",
			Content: "Content 2",
		},
	}, nil
}
