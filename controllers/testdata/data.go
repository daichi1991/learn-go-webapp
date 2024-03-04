package testdata

import "github.com/daichi1991/learn-go-webapp/models"

var articleTestData = []models.Article{
	models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "This is first post",
		UserName: "daichi",
		NiceNum: 2,
		CommentList: commentTestData,
	},
	models.Article{
		ID: 2,
		Title: "secondPost",
		Contents: "This is second post",
		UserName: "daichi",
		NiceNum: 4,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "This is first comment",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "This is second comment",
	},
}
