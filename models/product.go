package models

import (
	"mime/multipart"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryID  int       `json:"category_id"`
}

type ProductBody struct {
	ID          int                     `form:"id"`
	Name        string                  `form:"name"`
	Price       float64                 `form:"price"`
	Thumbnail   string                  `form:"thumbnail"`
	Description string                  `form:"description"`
	CreatedAt   *time.Time              `form:"created_at"`
	UpdatedAt   *time.Time              `form:"updated_at"`
	CategoryID  int                     `form:"category_id"`
	Files       []*multipart.FileHeader `form:"files"`
}
