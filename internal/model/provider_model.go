package model

import "time"

type Provider struct {
	ID        uint       `json:"id"`
	UUID      *string    `json:"uuid"`
	Code      *string    `json:"code" binding:"required"`
	Name      *string    `json:"name" binding:"required"`
	Address   *string    `json:"address" binding:"required"`
	Phone     *string    `json:"phone" binding:"required"`
	City      *string    `json:"city" binding:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
