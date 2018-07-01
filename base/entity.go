package base

import (
	"time"
)

//Model refers to the base entity for all other db model
type Model struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	CreateTime time.Time `json:"create_time" `
	UpdateTime time.Time `json:"update_time" `
}
