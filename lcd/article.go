package lcd

type Article struct {
	Uid               int64     `boil:"uid"`
	Owner             User      `boil:"owner"`
	Title             string    `boil:"title"`
	ConnectedArticles []Article `boil:"connected_articles"`
}
