package services

import(
	"github.com/daichi1991/learn-go-webapp/repositories"
	"github.com/daichi1991/learn-go-webapp/models"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}
