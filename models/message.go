package models

import "gorm.io/gorm"

type Message struct {
	gorm.models
	Content string `json:"content" gorm: "text; not null; default:null`
	Author string `json:"author" gorm: "text; not null; default:null`
}