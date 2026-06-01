package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name      string  `json:"name" gorm:"not null"`
    SKU       string  `json:"sku" gorm:"unique;not null"`
    Category  string  `json:"category"`
    Quantity  int     `json:"quantity" gorm:"default:0"`
    Price     float64 `json:"price"`
    Threshold int     `json:"threshold" gorm:"default:10"` // for low stock alert
}