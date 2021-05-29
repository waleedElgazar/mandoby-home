package db

import (
	"fmt"
	"time"
)

type Post struct {
	Phone      string
	Name       string
	Type       string
	Amount     string
	Government string
	Date       time.Time
}

func ToString(post Post) string {
	return fmt.Sprintf("phone %s name %s type %s amount %s government %s ", post.Phone, post.Name, post.Type, post.Amount, post.Government)
}
