package dao

import "gorm.io/gorm"

type Recipient struct {
	gorm.Model
	Name        string `gorm:"name;"`
	PhoneNumber string `gorm:"uniqueIndex;size:16"`
	CampaignID  string `gorm:"campaign_id;index"`
}

func (r *Recipient) TableName() string {
	return "tb_recipient"
}
