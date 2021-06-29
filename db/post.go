package db

import (
	"fmt"
)

type Post struct {
	PostID 		int `json:postId`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	ProductType string `json:"productType"`
	Amount      string `json:"amount"`
	Government  string `json:"government"`
	UserType    string `json:"userType"`
	Area        string `json:"area"`

	Date string
}

func ToString(post Post) string {
	return fmt.Sprintf(" phone %s name %s type %s amount %s government %s ",post.Phone, post.Name, post.ProductType, post.Amount, post.Government)
}
