package lcd

import (
	"database/sql"
	errors2 "errors"

	"github.com/ariyn/Lcd/util"
)

var (
	noRowsErr = errors2.New("sql: no rows in result set")
)

type MysqlUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		db: db,
	}
}

func (r MysqlUserRepository) GetUser(userId string) (u User, err error) {
	u.Id = userId
	err = r.db.QueryRow("SELECT `uid`, `name` FROM user WHERE id = ?", userId).Scan(&u.Uid, &u.Name)
	if err == noRowsErr {
		err = util.NoResultErr
	}
	return
}

func (r MysqlUserRepository) GetUserByUid(uid int64) (u User, err error) {
	u.Uid = uid
	err = r.db.QueryRow("SELECT `id`, `name` FROM user WHERE uid = ?", uid).Scan(&u.Id, &u.Name)
	if err == noRowsErr {
		err = util.NoResultErr
	}
	return
}

func (r MysqlUserRepository) CreateUser(userId, name string) (u User, err error) {
	u.Name = name
	u.Id = userId

	result, err := r.db.Exec("INSERT INTO user SET id=?, name=?", userId, name)
	if err != nil {
		return
	}

	u.Uid, err = result.LastInsertId()
	return
}

type MysqlArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *MysqlArticleRepository {
	return &MysqlArticleRepository{
		db: db,
	}
}

func (r MysqlArticleRepository) GetArticleByUid(uid int64) (a Article, err error) {
	a.Uid = uid
	err = r.db.QueryRow("SELECT `owner_uid`, `title` FROM article WHERE uid = ?", uid).Scan(&a.Owner.Uid, &a.Title)
	if err == noRowsErr {
		err = util.NoResultErr
	}
	return
}

func (r MysqlArticleRepository) CreateArticle(owner User, title string) (a Article, err error) {
	a.Owner = owner
	a.Title = title
	result, err := r.db.Exec("INSERT INTO article SET owner_uid=?, title=?", owner.Uid, title)
	if err != nil {
		return
	}

	a.Uid, err = result.LastInsertId()
	return
}
