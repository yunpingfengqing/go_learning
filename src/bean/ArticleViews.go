package bean

type ArticleViews struct {
	Date       string
	Article_Id int
	Views      int
}

func (ArticleViews) TableName() string {
	return "Article_Views"
}
