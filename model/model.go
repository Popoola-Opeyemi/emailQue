package model

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type EMailMsg struct {
	To      string `json:"to,omitempty"`
	HTML    string `json:"html,omitempty"`
	Text    string `json:"text,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type EmailtaskQueue struct {
	ID          int64           `json:"id"`
	Type        int64           `json:"type"`
	Data        json.RawMessage `json:"data"`
	Status      int64           `json:"status"`
	DateCreated time.Time       `json:"date_created"`
}

// Create ...
func (s *EMailMsg) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

// Update ...
func (s *EMailMsg) Update(db *gorm.DB, id int) error {
	return db.Model(&s).Where("id = ?", id).Update(s).Error
}

// Delete ...
func (s *EMailMsg) Delete(db *gorm.DB) error {
	return nil
}
