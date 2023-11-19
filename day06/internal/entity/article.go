package entity

type Article struct {
	ID      int
	Title   string `gorm:"type:text;notnull"`
	Content string `gorm:"type:text;notnull"`
}
