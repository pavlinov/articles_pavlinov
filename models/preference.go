package models

import (
    "github.com/jinzhu/gorm"
)

type Preference struct {
    gorm.Model
    UserID  uint `json:"user_id"`
    ArticleID uint `json:"article_id"`
    Liked   bool `json:"liked"`
}

