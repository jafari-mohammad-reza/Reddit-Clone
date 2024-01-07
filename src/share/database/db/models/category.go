package models

type Category struct {
	Name           string
	ParentCategory *Category
}
