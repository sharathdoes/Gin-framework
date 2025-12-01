package models

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
	Tags  pq.StringArray  `gorm:"type:text[]"`
	Comments []Comment  `gorm:"foreignKey:PostID"`
}
type Comment struct {
	gorm.Model
	PostID    uint     
	Message   string
	Author    string
}

