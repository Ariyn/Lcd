package lcd

type UserRepository interface {
	GetUser(userId string) (u User, err error)
	GetUserByUid(uid int64) (u User, err error)
	CreateUser(userId, name string) (u User, err error)
}

type ArticleRepository interface {
	GetArticleByUid(uid int64) (a Article, err error)
	CreateArticle(owner User, title string) (a Article, err error)
}
