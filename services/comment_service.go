package services

import(
	"github.com/daichi1991/learn-go-webapp/repositories"
	"github.com/daichi1991/learn-go-webapp/models"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()
	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}
