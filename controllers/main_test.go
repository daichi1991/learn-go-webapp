package controllers_test

import (
	"testing"
	"github.com/daichi1991/learn-go-webapp/controllers"
	"github.com/daichi1991/learn-go-webapp/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
