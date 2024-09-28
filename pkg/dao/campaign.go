package dao

import (
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	CampaignID   string `gorm:"campaign_id;unique_index"`
	Template     string `gorm:"template"`
	SchedualTime int64  `gorm:"schedual_time;index"`
}

func (c *Campaign) TableName() string {
	return "tb_campaign"
}
