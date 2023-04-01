package models

type Category struct {
	Id int64 `json:"id" db:"id"`
	CategoryName string `json:"category_name" db:"category_name"`
	category_no int8 `json:"category_no" db:"category_no"`
}