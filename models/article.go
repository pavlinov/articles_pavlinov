package models

import (
    "github.com/jinzhu/gorm"
)

type Article struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}

