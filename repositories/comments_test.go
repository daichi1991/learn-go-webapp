package repositories_test

import (
	"testing"
	"github.com/daichi1991/learn-go-webapp/models"
	"github.com/daichi1991/learn-go-webapp/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommntList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 	1,
		Message:		"test comment",
	}
	expectedCommentNum := 3
	_, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	sqlStr := `
		select count(*)
		from comments
		where article_id = ?;
	`
	var commentCount int
	row := testDB.QueryRow(sqlStr, comment.ArticleID)
	if err := row.Scan(&commentCount); err != nil {
		t.Fatal(err)
	}
	if commentCount != expectedCommentNum {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentNum, commentCount)
	}

	t.Cleanup(func() {
		sqlStr := `
			delete from comments
			where article_id = ?
			and message = ?;
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}
