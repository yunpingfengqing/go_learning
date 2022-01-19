package bean

type ArticleType struct {
	Id   int `gorm:"default:uuid_generate_v3(); PRIMARY_KEY"`
	Name string
}

func (ArticleType) TableName() string {
	return "Article_Type"
}
