package repositories

import(
	"database/sql"
	"github.com/daichi1991/learn-go-webapp/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
	`

	var newArtice models.Article
	newArtice.Title, newArtice.Contents, newArtice.UserName = article.Title, article.Contents, article.Username

	result, err := db.Exec(sqlStr, article.title, article.contents, article.username)
	if err != nil {
		return models.Article{}, err
	}
	id, _ := result.LastInsertId()
	newArtice.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`
	rows, err := db.Query(sqlStr, 5, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]modles.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.createdTime  )
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
