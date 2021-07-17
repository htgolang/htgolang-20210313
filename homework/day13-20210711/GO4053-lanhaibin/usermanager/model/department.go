package model

type Department struct {
	Id   int
	Name string `orm:"size(255)"`
}
