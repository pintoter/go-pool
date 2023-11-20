package entity

type Article struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"type:text;notnull"`
	Content string `gorm:"type:text;notnull"`
}
