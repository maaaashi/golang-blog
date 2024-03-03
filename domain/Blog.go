package entity

type Blog struct {
	Id      BlogId
	Title   BlogTItle
	Content BlogContent
}

type BlogId int
type BlogTItle string
type BlogContent string
