package models

import (
	"time"
)

type Comment struct { 
	CommentID 	int				`json:"comment_id"`
	ArticleID 	int				`json:"article_id"`
	Message 		string		`json:"message"`
	CreatedAt 	time.Time `json:"created_at"`
}

type Article struct {
	ID 					int				`json:"article_id"`
	Title 			string		`json:"title"`
	Contents 		string		`json:"content`
	UserName 		string		`json:"user_name"`
	NiceNum 		int				`json:"nice_num"`
	CommentList []Comment	`json:"comment_list"`
	CreatedAt 	time.Time	`json:"created_at"`
}
