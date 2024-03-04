package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/daichi1991/learn-go-webapp/models"
	"github.com/daichi1991/learn-go-webapp/controllers/services"
	"github.com/daichi1991/learn-go-webapp/apperrors"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHander(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHander(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}