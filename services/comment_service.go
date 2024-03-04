package services

import(
	"github.com/daichi1991/learn-go-webapp/repositories"
	"github.com/daichi1991/learn-go-webapp/models"
	"github.com/daichi1991/learn-go-webapp/apperrors"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}
